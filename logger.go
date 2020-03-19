package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/valyala/bytebufferpool"
)

type Logger struct {
	level Level
	out   io.Writer
}

func NewLogger(level Level, out io.Writer) *Logger {
	return &Logger{level: level, out: out}
}

func (l *Logger) output(level Level, s string) error {
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
	buf.WriteString(s)
	buf.WriteByte('\n')

	_, err := buf.WriteTo(l.out)
	bytebufferpool.Put(buf)
	return err
}

func (l *Logger) log(level Level, v ...interface{}) {
	if level <= l.level {
		l.output(level, fmt.Sprint(v...))
	}
}

func (l *Logger) logf(level Level, format string, v ...interface{}) {
	if level <= l.level {
		l.output(level, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	if PANIC <= l.level {
		l.output(PANIC, s)
	}
	panic(s)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	if PANIC <= l.level {
		l.output(PANIC, s)
	}
	panic(s)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.log(FATAL, v...)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logf(FATAL, format, v...)
	os.Exit(1)
}

func (l *Logger) Error(v ...interface{}) {
	l.log(ERROR, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logf(ERROR, format, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.log(WARN, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logf(WARN, format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.log(INFO, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logf(INFO, format, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.log(DEBUG, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logf(DEBUG, format, v...)
}

func (l *Logger) Trace(v ...interface{}) {
	l.log(TRACE, v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.logf(TRACE, format, v...)
}
