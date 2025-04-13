package bunjang

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	Query(ctx context.Context) ([]Product, error)
}

func NewClientWithConfig(c ClientConfig) (Client, error) {
	return &client{
		config: c,
		client: &http.Client{},
	}, nil
}

type client struct {
	config ClientConfig
	client *http.Client
}

type ClientConfig struct {
	URL     *url.URL
	Query   string
	Timeout time.Duration
}

func (c *client) Query(ctx context.Context) ([]Product, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", c.config.URL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "bunjang-watch/1.0 (+https://github.com/starpia-forge/bunjang-watch)")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp Response
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	return apiResp.List, nil
}
