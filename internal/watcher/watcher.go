package watcher

import (
	"context"
	"github.com/starpia-forge/bunjang-watch/internal/watcher/api/v1"
	"time"
)

type Watcher interface {
	Watch(ctx context.Context) (chan []v1.Product, error)
}

type watcher struct {
	*WatcherConfig
}

func NewWatcher() Watcher {
	w := &watcher{
		WatcherConfig: &DefaultWatcherConfig,
	}
	return w
}

func NewWatcherWithConfig(c *WatcherConfig) Watcher {
	w := &watcher{
		WatcherConfig: c,
	}
	return w
}

func (w *watcher) Watch(ctx context.Context) (chan []v1.Product, error) {
	out := make(chan []v1.Product)

	go func() {
		ticker := time.NewTicker(w.Interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				// TODO - Error Handling
				return
			case <-ticker.C:
			}
			products, err := w.watch(ctx)
			if err != nil {
				// TODO - Error Handling
			}
			out <- products
		}
	}()

	return out, nil
}

func (w *watcher) watch(ctx context.Context) ([]v1.Product, error) {
	products, err := v1.Query(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
