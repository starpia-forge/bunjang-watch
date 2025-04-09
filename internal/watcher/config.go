package watcher

import (
	"regexp"
	"time"
)

var DefaultWatcherConfig WatcherConfig

type WatcherConfig struct {
	Query          string
	Keywords       []*regexp.Regexp
	IgnoreKeywords []*regexp.Regexp
	MinPrice       int
	MaxPrice       int
	IncludeUsed    bool
	Interval       time.Duration
}
