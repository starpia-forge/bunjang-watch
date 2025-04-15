package watcher

import (
	"context"
	"github.com/starpia-forge/bunjang-watch/internal/bunjang"
	"github.com/starpia-forge/bunjang-watch/internal/notifier"
	"github.com/starpia-forge/bunjang-watch/internal/watcher/filter"
	"time"
)

type Watcher interface {
	Watch(ctx context.Context) (chan []bunjang.Product, error)
}

type WatcherOptions func(w *watcher)

func WithWatcherFilters(f ...filter.Filter[bunjang.Product]) WatcherOptions {
	return func(w *watcher) {
		w.filters = f
	}
}

func WithWatcherClient(client bunjang.Client) WatcherOptions {
	return func(w *watcher) {
		w.client = client
	}
}

func WithWatcherNotifier(n notifier.Notifier) WatcherOptions {
	return func(w *watcher) {
		w.notifier = n
	}
}

func NewWatcher(c WatcherConfig, opts ...WatcherOptions) Watcher {
	w := &watcher{
		config: c,
	}

	for _, opt := range opts {
		opt(w)
	}

	return w
}

type watcher struct {
	config   WatcherConfig
	filters  []filter.Filter[bunjang.Product]
	client   bunjang.Client
	notifier notifier.Notifier
}

func (w *watcher) Watch(ctx context.Context) (chan []bunjang.Product, error) {
	out := make(chan []bunjang.Product)

	go func() {
		defer close(out)
		ticker := time.NewTicker(w.config.Interval)
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

func (w *watcher) watch(ctx context.Context) ([]bunjang.Product, error) {
	products, err := w.client.Query(ctx)
	if err != nil {
		return nil, err
	}
	return w.filter(products), nil
}

func (w *watcher) filter(products []bunjang.Product) []bunjang.Product {
	var result []bunjang.Product
	for _, product := range products {
		apply := true
		for _, productFilter := range w.filters {
			if apply = productFilter.Apply(product); !apply {
				break
			}
		}
		if apply {
			result = append(result, product)
		}
	}
	return result
}
