package hook

import (
	"context"
	"net/url"
)

type slackHook struct {
	config HookConfig
	url    *url.URL
}

func newSlackHook(c HookConfig) (*slackHook, error) {
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}
	return &slackHook{
		config: c,
		url:    u,
	}, nil
}

func (s *slackHook) Type() string {
	return s.config.Type
}

func (s *slackHook) Name() string {
	return s.config.Name
}

func (s *slackHook) SendAlert(ctx context.Context, alert string) error {
	return nil
}

func (s *slackHook) setURL(u *url.URL) {
	s.url = u
}
