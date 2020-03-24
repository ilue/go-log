package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/valyala/bytebufferpool"
)

type Logger struct {
	name  string
	level Level
	out   io.Writer
}

func NewLogger(name string, level Level, out io.Writer) *Logger {
	return &Logger{name: name, level: level, out: out}
}

func (l *Logger) WithName(name string) *Logger {
	return NewLogger(name, l.level, l.out)
}

func (l *Logger) WithOutput(out io.Writer) *Logger {
	return NewLogger(l.name, l.level, out)
}

func (l *Logger) enabled(level Level) bool {
	return l.level >= level
}

func (l *Logger) log(level Level, s string) error {
	now := time.Now()
	buf := bytebufferpool.Get()

	{
		year, month, day := now.Date()
		hour, min, sec := now.Clock()
		msec := now.Nanosecond() / 1e6
		itoa(buf, year, 4)
		buf.WriteByte('-')
		itoa(buf, int(month), 2)
		buf.WriteByte('-')
		itoa(buf, day, 2)
		buf.WriteByte(' ')
		itoa(buf, hour, 2)
		buf.WriteByte(':')
		itoa(buf, min, 2)
		buf.WriteByte(':')
		itoa(buf, sec, 2)
		buf.WriteByte('.')
		itoa(buf, msec, 3)
	}
	buf.WriteByte(' ')
	{
		levelStr := _levelTexts[level]
		buf.WriteString(levelStr)
		const padding = " "
		buf.WriteString(padding[:5-len(levelStr)])
	}
	buf.WriteByte(' ')
	if l.name != "" {
		buf.WriteString(l.name)
		buf.WriteString(" - ")
	}
	buf.WriteString(s)
	buf.WriteByte('\n')

	_, err := buf.WriteTo(l.out)
	bytebufferpool.Put(buf)
	return err
}

func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	if l.enabled(PANIC) {
		l.log(PANIC, s)
	}
	panic(s)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	if l.enabled(PANIC) {
		l.log(PANIC, s)
	}
	panic(s)
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.enabled(FATAL) {
		l.log(FATAL, fmt.Sprint(v...))
	}
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.enabled(FATAL) {
		l.log(FATAL, fmt.Sprintf(format, v...))
	}
	os.Exit(1)
}

func (l *Logger) Error(v ...interface{}) {
	if l.enabled(ERROR) {
		l.log(ERROR, fmt.Sprint(v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.enabled(ERROR) {
		l.log(ERROR, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if l.enabled(WARN) {
		l.log(WARN, fmt.Sprint(v...))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.enabled(WARN) {
		l.log(WARN, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.enabled(INFO) {
		l.log(INFO, fmt.Sprint(v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.enabled(INFO) {
		l.log(INFO, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if l.enabled(DEBUG) {
		l.log(DEBUG, fmt.Sprint(v...))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.enabled(DEBUG) {
		l.log(DEBUG, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Trace(v ...interface{}) {
	if l.enabled(TRACE) {
		l.log(TRACE, fmt.Sprint(v...))
	}
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.enabled(TRACE) {
		l.log(TRACE, fmt.Sprintf(format, v...))
	}
}
