package log

import (
	"context"
)

type contextKey string

const (
	ContextKeyLogger contextKey = "__logger__"
)

func Context(ctx context.Context, l Logger, with ...interface{}) context.Context {
	if l == nil {
		l = Current()
	}
	return context.WithValue(ctx, ContextKeyLogger, l.With(with...))
}

func C(ctx context.Context) (l Logger) {
	var ok bool
	if loggerItf := ctx.Value(ContextKeyLogger); loggerItf == nil {
		return Current()
	} else if l, ok = loggerItf.(Logger); !ok {
		return Current()
	} else {
		return l
	}
}
