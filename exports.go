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
	AddDepth(depth int) Logger
	StdLogger() *stdLog.Logger

	Sync()
}

var logger = NewSimpleLogger("", os.Stdout, nil).AddDepth(1)

func Use(l Logger) {
	if logger != nil {
		var oldLogger = logger
		defer func() {
			oldLogger.Sync()
		}()
	}
	l = l.AddDepth(1)
	logger = l
}

func Current() Logger {
	var l = logger
	l = l.AddDepth(-1)
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
	return Current().With(v...)
}

func WithName(name string) Logger {
	return Current().WithName(name)
}

func Sync() {
	logger.Sync()
}

func StdLogger() *stdLog.Logger {
	return Current().StdLogger()
}
