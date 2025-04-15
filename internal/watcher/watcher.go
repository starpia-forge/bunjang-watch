package watcher

import (
	"context"
	"github.com/starpia-forge/bunjang-watch/internal/bunjang"
	"github.com/starpia-forge/bunjang-watch/internal/filter"
	"time"
)

type Watcher interface {
	Watch(ctx context.Context) (chan []bunjang.Product, error)
}

func NewWatcher(c WatcherConfig) Watcher {
	w := &watcher{
		config: c,
	}
	return w
}

type watcher struct {
	config         WatcherConfig
	productFilters []filter.Filter[bunjang.Product]
	client         bunjang.Client
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
		for _, productFilter := range w.productFilters {
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
