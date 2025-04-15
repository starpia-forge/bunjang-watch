package bunjang

import (
	"net/url"
	"time"
)

type ClientConfig struct {
	URL     *url.URL
	Query   string
	Timeout time.Duration
}
