package watcher

import (
	"context"
	"github.com/starpia-forge/bunjang-watch/internal/watcher/api"
)

type Watcher interface {
	FirstRun() bool
	Do(ctx context.Context) ([]api.Product, error)
}

type watcher struct {
	firstRun bool
	products map[string]api.Product
	config   Config
}

func NewWatcher(c Config) Watcher {
	return &watcher{
		firstRun: true,
		config:   c,
	}
}

func (w *watcher) FirstRun() bool {
	return w.firstRun
}

func (w *watcher) Do(context.Context) ([]api.Product, error) {
	return []api.Product{}, nil
}
