package watcher

import (
	"regexp"
	"time"
)

var defaultWatchConfig WatchConfig

type WatchConfig struct {
	Interval    time.Duration
	Keywords    []*regexp.Regexp
	MinPrice    int
	MaxPrice    int
	IncludeUsed bool
}
