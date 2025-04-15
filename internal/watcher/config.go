package watcher

import (
	"github.com/starpia-forge/bunjang-watch/internal/watcher/filter"
	"time"
)

var DefaultWatcherConfig WatcherConfig

type WatcherConfig struct {
	Query    string              `json:"query,omitempty"`
	Interval time.Duration       `json:"interval,omitempty"`
	Filter   filter.FilterConfig `json:"filter"`
}
