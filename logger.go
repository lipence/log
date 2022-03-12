package log

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	stdLog "log"
	"os"
	"strings"
)

func toString(value interface{}) (str string) {
	switch _value := value.(type) {
	case []byte:
		value = "0x" + hex.EncodeToString(_value)
	case fmt.Stringer:
		value = _value.String()
	case fmt.GoStringer:
		value = _value.GoString()
	}
	if jBytes, err := json.Marshal(value); err == nil {
		str = string(jBytes)
	} else {
		str = fmt.Sprintln(value)
	}
	return str
}

func NewSimpleLogger(name string, output io.Writer, data [][2]string) *simple {
	var prefix string
	if name != "" {
		prefix = "[" + name + "]"
	}
	var sysLogger = stdLog.New(
		output, prefix, stdLog.Ldate|stdLog.Ltime|stdLog.Lmicroseconds|stdLog.Lshortfile|stdLog.Lmsgprefix,
	)
	_logger := &simple{name: name, dataArr: data, logger: sysLogger}
	_logger.buildDataStr()
	return _logger
}

type simple struct {
	name    string
	dataArr [][2]string
	dataStr string
	depth   int
	logger  *stdLog.Logger
}

func (s *simple) Debug(v ...interface{}) {
	s.errHandler(s.logger.Output(s.depth+2, "[D] "+fmt.Sprint(v...)+s.dataStr))
}

func (s *simple) Debugf(format string, v ...interface{}) {
	s.errHandler(s.logger.Output(s.depth+2, "[D] "+fmt.Sprintf(format, v...)+s.dataStr))
}

func (s *simple) Info(v ...interface{}) {
	s.errHandler(s.logger.Output(s.depth+2, "[I] "+fmt.Sprint(v...)+s.dataStr))
}

func (s *simple) Infof(format string, v ...interface{}) {
	s.errHandler(s.logger.Output(s.depth+2, "[I] "+fmt.Sprintf(format, v...)+s.dataStr))
}

func (s *simple) Warn(v ...interface{}) {
	s.errHandler(s.logger.Output(s.depth+2, "[W] "+fmt.Sprint(v...)+s.dataStr))
}

func (s *simple) Warnf(format string, v ...interface{}) {
	s.errHandler(s.logger.Output(s.depth+2, "[W] "+fmt.Sprintf(format, v...)+s.dataStr))
}

func (s *simple) Error(v ...interface{}) {
	s.errHandler(s.logger.Output(s.depth+2, "[E] "+fmt.Sprint(v...)+s.dataStr))
}

func (s *simple) Errorf(format string, v ...interface{}) {
	s.errHandler(s.logger.Output(s.depth+2, "[E] "+fmt.Sprintf(format, v...)+s.dataStr))
}

func (s *simple) Fatal(v ...interface{}) {
	content := fmt.Sprint(v...)
	s.errHandler(s.logger.Output(s.depth+2, "[F] "+content+s.dataStr))
	s.Sync()
	os.Exit(-1)
}

func (s *simple) Fatalf(format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	s.errHandler(s.logger.Output(s.depth+2, "[F] "+content+s.dataStr))
	s.Sync()
	os.Exit(-1)
}

func (s *simple) Panic(v ...interface{}) {
	content := fmt.Sprint(v...)
	s.errHandler(s.logger.Output(s.depth+2, "[P] "+content+s.dataStr))
	s.Sync()
	panic(content)
}

func (s *simple) Panicf(format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	s.errHandler(s.logger.Output(s.depth+2, "[P] "+content+s.dataStr))
	s.Sync()
	panic(content)
}

func (s *simple) Print(v ...interface{}) {
	s.errHandler(s.logger.Output(s.depth+2, "[I] "+fmt.Sprint(v...)+s.dataStr))
}

func (s *simple) Printf(format string, v ...interface{}) {
	s.errHandler(s.logger.Output(s.depth+2, "[I] "+fmt.Sprintf(format, v...)+s.dataStr))
}

func (s *simple) With(v ...interface{}) Logger {
	if len(v)%2 > 0 {
		v = append(v, nil)
	}
	var newData = make([][2]string, len(v)/2)
	for i := 0; i < len(v); i += 2 {
		newData[i/2] = [2]string{toString(v[i]), toString(v[i+1])}
	}
	return NewSimpleLogger(s.name, s.logger.Writer(), append(s.dataArr, newData...))
}

func (s *simple) WithName(name string) Logger {
	return NewSimpleLogger(name, s.logger.Writer(), s.dataArr)
}

func (s *simple) AddDepth(depth int) Logger {
	return &simple{
		name:    s.name,
		dataArr: s.dataArr,
		dataStr: s.dataStr,
		depth:   s.depth + depth,
		logger:  s.logger,
	}
}

func (s *simple) buildDataStr() {
	if len(s.dataArr) == 0 {
		s.dataStr = ""
		return
	}
	var builder = strings.Builder{}
	builder.Write([]byte{0x20, '/', '/', 0x20})
	for i, item := range s.dataArr {
		builder.WriteString(item[0])
		builder.WriteRune(':')
		builder.WriteString(item[1])
		if i+1 < len(s.dataArr) {
			builder.Write([]byte{',', 0x20})
		}
	}
	s.dataStr = builder.String()
}

func (s *simple) StdLogger() *stdLog.Logger {
	return s.logger
}

type flusher interface {
	Flush() error
}

func (s *simple) Sync() {
	if _flusher, ok := s.logger.Writer().(flusher); ok {
		if err := _flusher.Flush(); err != nil {
			fmt.Printf("logger sync filed: %v\n", err)
		}
	}
}

func (s *simple) errHandler(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "log output fail: %v", err)
	}
}
