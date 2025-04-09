package watcher

import (
	"context"
	"github.com/starpia-forge/bunjang-watch/internal/watcher/api"
)

type Watcher interface {
	Watch(ctx context.Context) ([]api.Product, error)
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

func (w *watcher) Watch(context.Context) ([]api.Product, error) {
	return []api.Product{}, nil
}
