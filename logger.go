package log

import (
	"io"
	"os"
	"time"

	"github.com/valyala/bytebufferpool"
)

type Logger struct {
	level   Level
	out     io.Writer
	context []byte
}

func defaultLogger() Logger {
	return Logger{level: DebugLevel, out: os.Stderr}
}

var Default = defaultLogger()

func With() *LoggerBuilder {
	return &LoggerBuilder{logger: Default}
}

func (l *Logger) With() *LoggerBuilder {
	return &LoggerBuilder{logger: *l}
}

func Panic(msg string, args ...interface{}) {
	Default.log(PanicLevel, msg, args)
	panic(msg)
}

func (l *Logger) Panic(msg string, args ...interface{}) {
	l.log(PanicLevel, msg, args)
	panic(msg)
}

func Fatal(msg string, args ...interface{}) {
	Default.log(FatalLevel, msg, args)
	os.Exit(1)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.log(FatalLevel, msg, args)
	os.Exit(1)
}

func Error(msg string, args ...interface{}) {
	Default.log(ErrorLevel, msg, args)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.log(ErrorLevel, msg, args)
}

func Warn(msg string, args ...interface{}) {
	Default.log(WarnLevel, msg, args)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.log(WarnLevel, msg, args)
}

func Info(msg string, args ...interface{}) {
	Default.log(InfoLevel, msg, args)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.log(InfoLevel, msg, args)
}

func Debug(msg string, args ...interface{}) {
	Default.log(DebugLevel, msg, args)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log(DebugLevel, msg, args)
}

func (l *Logger) log(level Level, msg string, args []interface{}) {
	if level > l.level {
		return
	}
	bb := bytebufferpool.Get()
	defer bytebufferpool.Put(bb)
	bb.B = l.formatHeader(bb.B, level)
	bb.B = append(bb.B, msg...)
	bb.B = append(bb.B, l.context...)
	bb.B = appendFields(bb.B, args)
	bb.B = append(bb.B, '\n')
	l.out.Write(bb.B)
}

func (l *Logger) formatHeader(buf []byte, level Level) []byte {
	{
		now := time.Now()
		_, month, day := now.Date()
		hour, min, sec := now.Clock()
		buf = itoa(buf, int(month), 2)
		buf = append(buf, '-')
		buf = itoa(buf, day, 2)
		buf = append(buf, ' ')
		buf = itoa(buf, hour, 2)
		buf = append(buf, ':')
		buf = itoa(buf, min, 2)
		buf = append(buf, ':')
		buf = itoa(buf, sec, 2)
		buf = append(buf, '.')
		buf = itoa(buf, now.Nanosecond()/1e6, 3)
	}
	buf = append(buf, ' ')
	{
		s := _levelText[level]
		buf = append(buf, s...)
		buf = append(buf, " "[:5-len(s)]...)
	}
	return append(buf, ' ')
}
