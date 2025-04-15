package hook

import (
	"context"
	"net/url"
)

type nateOnHook struct {
	config HookConfig
	url    *url.URL
}

func newNateOnHook(c HookConfig) (*nateOnHook, error) {
	return &nateOnHook{
		config: c,
	}, nil
}

func (n *nateOnHook) Type() string {
	return n.config.Type
}

func (n *nateOnHook) Name() string {
	return n.config.Name
}

func (n *nateOnHook) SendAlert(ctx context.Context, alert string) error {
	return nil
}

func (n *nateOnHook) setURL(u *url.URL) {
	n.url = u
}
