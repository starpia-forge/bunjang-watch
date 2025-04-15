package hook

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
	ctx, cancel := context.WithTimeout(ctx, n.config.Timeout)
	defer cancel()

	v := url.Values{}
	v.Set("content", alert)

	req, err := http.NewRequestWithContext(ctx, "POST", n.url.String(), strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("%w: %d", ErrRequestNotSuccessful, resp.StatusCode)
	}

	return nil
}

func (n *nateOnHook) setURL(u *url.URL) {
	n.url = u
}
