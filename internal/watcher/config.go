package watcher

import "time"

type Config struct {
	Interval time.Duration `json:"interval"`
	Filter
}

type Filter struct {
	Keywords     []string
	MinimumPrice int
	MaximumPrice int
	IncludeUsed  bool
}
