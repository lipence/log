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
		l = logger
	}
	return context.WithValue(ctx, ContextKeyLogger, l.With(with...))
}

func C(ctx context.Context) (l Logger) {
	var ok bool
	if loggerItf := ctx.Value(ContextKeyLogger); loggerItf == nil {
		return logger
	} else if l, ok = loggerItf.(Logger); !ok {
		return logger
	} else {
		return l
	}
}
