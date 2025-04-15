package notifier

import (
	"context"
	"github.com/starpia-forge/bunjang-watch/internal/hook"
)

type Notifier interface {
	Notify(context.Context, string) error
}

type MultiHookNotifier struct {
	hooks []hook.Hook
}

func (m *MultiHookNotifier) Notify(ctx context.Context, alert string) error {
	for _, h := range m.hooks {
		if err := h.SendAlert(ctx, alert); err != nil {
			return err
		}
	}
	return nil
}
