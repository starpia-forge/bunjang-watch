package watcher

import (
	"regexp"
	"time"
)

var DefaultWatcherConfig WatcherConfig

type WatcherConfig struct {
	Interval    time.Duration
	Keywords    []*regexp.Regexp
	MinPrice    int
	MaxPrice    int
	IncludeUsed bool
}
