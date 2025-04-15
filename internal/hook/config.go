package hook

import "time"

type HookConfig struct {
	Type    string        `json:"type,omitempty"`
	Name    string        `json:"name,omitempty"`
	URL     string        `json:"url,omitempty"`
	Timeout time.Duration `json:"timeout,omitempty"`
}
