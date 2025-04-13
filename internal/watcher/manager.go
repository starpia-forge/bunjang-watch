package watcher

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrWatcherAlreadyExists  = errors.New("watcher already exists with the given id")
	ErrWatcherNotFound       = errors.New("no watcher found with the given id")
	ErrWatcherAlreadyRunning = errors.New("watcher is already running")
	ErrWatcherNotRunning     = errors.New("watcher is not running")
)

type WatcherManager struct {
	mu       sync.RWMutex
	watchers map[string]*runningWatcher
}

type runningWatcher struct {
	watcher Watcher
	cancel  context.CancelFunc
	done    chan struct{}
}

func NewWatcherManager() *WatcherManager {
	return &WatcherManager{
		watchers: make(map[string]*runningWatcher),
	}
}

func (wm *WatcherManager) AddWatcher(id string, w Watcher) error {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	if _, exists := wm.watchers[id]; exists {
		return ErrWatcherAlreadyExists
	}

	wm.watchers[id] = &runningWatcher{
		watcher: w,
	}
	return nil
}

func (wm *WatcherManager) StartWatcher(id string) error {
	wm.mu.Lock()
	defer wm.mu.Unlock()

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
		for range out {
		}
	}()

	return nil
}

func (wm *WatcherManager) StopWatcher(id string) error {
	wm.mu.Lock()
	defer wm.mu.Unlock()

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

	wm.mu.Lock()
	defer wm.mu.Unlock()

	delete(wm.watchers, id)
	return nil
}

func (wm *WatcherManager) ListWatchers() []string {
	wm.mu.RLock()
	defer wm.mu.RUnlock()

	ids := make([]string, 0, len(wm.watchers))
	for id := range wm.watchers {
		ids = append(ids, id)
	}
	return ids
}
