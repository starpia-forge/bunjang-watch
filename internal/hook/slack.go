package hook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type slackHook struct {
	config HookConfig
	url    *url.URL
}

func newSlackHook(c HookConfig) (*slackHook, error) {
	return &slackHook{
		config: c,
	}, nil
}

func (s *slackHook) Type() string {
	return s.config.Type
}

func (s *slackHook) Name() string {
	return s.config.Name
}

func (s *slackHook) SendAlert(ctx context.Context, alert string) error {
	ctx, cancel := context.WithTimeout(ctx, s.config.Timeout)
	defer cancel()

	v := map[string]string{
		"text": alert,
	}
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", s.url.String(), bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

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

func (s *slackHook) setURL(u *url.URL) {
	s.url = u
}
