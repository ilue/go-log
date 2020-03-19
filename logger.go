package log

import (
	"fmt"
	"io"
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
	buf.WriteString(" ")
	{
		levelStr := _levelTexts[level]
		buf.WriteString(levelStr)
		const padding = " "
		buf.WriteString(padding[:5-len(levelStr)])
	}
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

func (l *Logger) Info(v ...interface{}) {
	l.log(INFO, v...)
}
