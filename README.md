# lipence/log

Package `log` is a simple wrap of native log package, which extract supports message level, logger name and additional data.

## Interface

```go
package log

import stdLog "log"

type Logger interface {
	Debug(v ...interface{}) // output as level Debug
	Debugf(format string, v ...interface{})

	Info(v ...interface{}) // output as level Info
	Infof(format string, v ...interface{})

	Warn(v ...interface{}) // output as level Warn
	Warnf(format string, v ...interface{})

	Error(v ...interface{}) // output as level Error
	Errorf(format string, v ...interface{})

	Fatal(v ...interface{}) // output as level Fatal (will cause os.exit)
	Fatalf(format string, v ...interface{})

	Panic(v ...interface{}) // output as level Panic (will cause panic)
	Panicf(format string, v ...interface{})

	Print(v ...interface{}) // alias of `Info` / `Infof`
	Printf(format string, v ...interface{})

	With(v ...interface{}) Logger // fork current logger and add data
	WithName(name string) Logger  // fork current logger and change name
	StdLogger() *stdLog.Logger    // get underlying std logger
	AddDepth(depth int) Logger    // set caller skip (offset)
	Sync()                        // write data
}

type WithSyncer interface { // optional for external logger
	Logger
	Sync() // synchronize logs to be output
}

type WithDepth interface { // optional for external logger
	Logger
	AddDepth(depth int) Logger // change callstack depth
}
```

## Usage

**OnInit**: package automatically create `*simple` instance as default logger

**BeforeStartup** (optional): replace default logger with external logger: e.g.

```
func main() {
	log.Use(anotherLoggerPkg.New())
}
```

**Running**: use logger: e.g.

```
func apiPutOrderHandler(ctx context.Context) {
	log.Info("start processing")
	// or with format
	log.Warnf("unexpected things happed")
	// or with some data
	log.With("transId", ctx.Value("transId")).Error("failed")
}
```

**NewTransaction** (optional): fork new logger instance and add data: e.g.

```
func apiPutOrderHandler(ctx context.Context) {
	logger := log.With("api", "putOrder")
	logger.Debug("start processing")
}
```

**BeforeShutdown**: sync log: e.g.

```
func BeforeShutdown() {
	log.Sync()
}
```

## Changelog

- v0.1.0 Initial Version

## Contact

Kenta Lee ( [kenta.li@cardinfolink.com](mailto:kenta@cardinfolink.com) )

## License

`lipence/log` source code is available under the Apache-2.0 [License](/LICENSE)
