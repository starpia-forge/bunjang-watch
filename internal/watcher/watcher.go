package watcher

type Watcher interface {
	WatchListings()
}

type watcher struct {
	config Config
}

func NewWatcher(c Config) Watcher {
	return &watcher{
		config: c,
	}
}

func (w *watcher) WatchListings() {}
