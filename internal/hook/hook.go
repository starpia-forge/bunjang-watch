package hook

import (
	"context"
	"errors"
	"net/url"
)

const (
	HookTypeSlack  = "slack"
	HookTypeNateOn = "nate_on"
)

var (
	ErrInvalidHookType = errors.New("invalid hook type")
)

type Hook interface {
	Type() string
	Name() string                            // 훅의 이름 (로깅/식별용)
	SendAlert(context.Context, string) error // 알람 전송

	setURL(*url.URL)
}

func NewHook(c HookConfig) (Hook, error) {
	switch c.Type {
	case HookTypeNateOn:
		return newHook[*nateOnHook](c, newNateOnHook)
	case HookTypeSlack:
		return newHook[*slackHook](c, newSlackHook)
	default:
	}
	return nil, ErrInvalidHookType
}

type hookConstructorFunc[T Hook] func(c HookConfig) (T, error)

func newHook[T Hook](c HookConfig, constructor hookConstructorFunc[T]) (T, error) {
	var hook T
	u, err := url.Parse(c.URL)
	if err != nil {
		return hook, err
	}

	hook, err = constructor(c)
	if err != nil {
		return hook, err
	}

	hook.setURL(u)

	return hook, nil
}
