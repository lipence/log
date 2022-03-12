package log

import (
	stdLog "log"
	"os"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	Panic(v ...interface{})
	Panicf(format string, v ...interface{})

	Print(v ...interface{})
	Printf(format string, v ...interface{})

	With(v ...interface{}) Logger
	WithName(name string) Logger
	StdLogger() *stdLog.Logger
}

type WithSyncer interface {
	Logger
	Sync()
}

type WithDepth interface {
	Logger
	AddDepth(depth int) Logger
}

var logger = NewSimpleLogger("", os.Stdout, nil).AddDepth(1)

func Use(l Logger) {
	if logger != nil {
		var oldLogger = logger
		defer func() {
			if _s, ok := oldLogger.(WithSyncer); ok {
				_s.Sync()
			}
		}()
	}
	if _d, ok := l.(WithDepth); ok {
		l = _d.AddDepth(1)
	}
	logger = l
}

func Current() Logger {
	var l = logger
	if _d, ok := l.(WithDepth); ok {
		l = _d.AddDepth(-1)
	}
	return l
}

func Debug(m ...interface{}) {
	logger.Debug(m...)
}

func Info(m ...interface{}) {
	logger.Info(m...)
}

func Warn(m ...interface{}) {
	logger.Warn(m...)
}

func Error(m ...interface{}) {
	logger.Error(m...)
}

func Panic(m ...interface{}) {
	logger.Panic(m...)
}

func Fatal(m ...interface{}) {
	logger.Fatal(m...)
}

func Print(m ...interface{}) {
	logger.Print(m...)
}

func Debugf(format string, m ...interface{}) {
	logger.Debugf(format, m...)
}

func Infof(format string, m ...interface{}) {
	logger.Infof(format, m...)
}

func Warnf(format string, m ...interface{}) {
	logger.Warnf(format, m...)
}

func Errorf(format string, m ...interface{}) {
	logger.Errorf(format, m...)
}

func Panicf(format string, m ...interface{}) {
	logger.Panicf(format, m...)
}

func Fatalf(format string, m ...interface{}) {
	logger.Fatalf(format, m...)
}

func Printf(format string, m ...interface{}) {
	logger.Infof(format, m...)
}

func With(v ...interface{}) Logger {
	return logger.With(v...)
}

func Sync() {
	if _s, ok := logger.(WithSyncer); ok {
		_s.Sync()
	}
}

func StdLogger() *stdLog.Logger {
	return logger.StdLogger()
}
