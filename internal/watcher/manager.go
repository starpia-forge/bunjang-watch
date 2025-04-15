package watcher

import (
	"context"
	"errors"
	"github.com/starpia-forge/bunjang-watch/internal/bunjang"
	"github.com/starpia-forge/bunjang-watch/internal/notifier"
	"github.com/starpia-forge/bunjang-watch/internal/watcher/filter"
	"sync"
)

var (
	ErrWatcherAlreadyExists  = errors.New("watcher already exists with the given id")
	ErrWatcherNotFound       = errors.New("no watcher found with the given id")
	ErrWatcherAlreadyRunning = errors.New("watcher is already running")
	ErrWatcherNotRunning     = errors.New("watcher is not running")
)

type WatcherManager struct {
	sync.RWMutex
	watchers map[string]*runningWatcher
}

type runningWatcher struct {
	watcher  Watcher
	filters  []filter.Filter[bunjang.Product]
	client   bunjang.Client
	notifier notifier.Notifier
	cancel   context.CancelFunc
	done     chan struct{}
}

func NewWatcherManager() *WatcherManager {
	return &WatcherManager{
		watchers: make(map[string]*runningWatcher),
	}
}

func (wm *WatcherManager) AddWatcher(id string, w WatcherConfig) error {
	wm.Lock()
	defer wm.Unlock()

	if _, exists := wm.watchers[id]; exists {
		return ErrWatcherAlreadyExists
	}

	wm.watchers[id] = &runningWatcher{
		watcher: NewWatcher(w),
	}
	return nil
}

func (wm *WatcherManager) StartWatcher(id string) error {
	wm.Lock()
	defer wm.Unlock()

	rw, ok := wm.watchers[id]
	if !ok {
		return ErrWatcherNotFound
	}

	if rw.cancel != nil {
		return ErrWatcherAlreadyRunning
	}

	ctx, cancel := context.WithCancel(context.Background())
	out, err := rw.watcher.Watch(ctx)
	if err != nil {
		cancel()
		return err
	}

	rw.cancel = cancel
	rw.done = make(chan struct{})

	go func() {
		defer close(rw.done)
		for {
			select {
			case <-ctx.Done():
				return
			case products := <-out:
				for _, product := range filter.ChainApply(rw.filters, products) {
					// TODO - Error Handling
					rw.notifier.Notify(ctx, product.Name)
				}
			}
		}
	}()

	return nil
}

func (wm *WatcherManager) StopWatcher(id string) error {
	wm.Lock()
	defer wm.Unlock()

	rw, ok := wm.watchers[id]
	if !ok {
		return ErrWatcherNotFound
	}

	if rw.cancel == nil {
		return ErrWatcherNotRunning
	}

	rw.cancel()
	<-rw.done
	rw.cancel = nil
	rw.done = nil

	return nil
}

func (wm *WatcherManager) RemoveWatcher(id string) error {
	if err := wm.StopWatcher(id); err != nil &&
		!errors.Is(err, context.Canceled) &&
		!errors.Is(err, ErrWatcherNotRunning) {
		return err
	}

	wm.Lock()
	defer wm.Unlock()

	delete(wm.watchers, id)
	return nil
}

func (wm *WatcherManager) ListWatchers() []string {
	wm.RLock()
	defer wm.RUnlock()

	ids := make([]string, 0, len(wm.watchers))
	for id := range wm.watchers {
		ids = append(ids, id)
	}
	return ids
}
