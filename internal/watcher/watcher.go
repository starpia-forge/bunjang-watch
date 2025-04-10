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

type watcher struct {
	*WatcherConfig
	productFilters []filter.Filter[bunjang.Product]
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

func (w *watcher) Watch(ctx context.Context) (chan []bunjang.Product, error) {
	out := make(chan []bunjang.Product)

	go func() {
		defer close(out)
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

func (w *watcher) watch(ctx context.Context) ([]bunjang.Product, error) {
	products, err := bunjang.Query(ctx)
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
