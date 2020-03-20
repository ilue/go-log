package log

import (
	"io"
	"os"
)

var Default = NewLogger(DEBUG, SyncWriter(os.Stderr))

func Output(w io.Writer) *Logger {
	return Default.Output(w)
}

func Panic(v ...interface{}) {
	Default.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	Default.Panicf(format, v...)
}

func Fatal(v ...interface{}) {
	Default.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	Default.Fatalf(format, v...)
}

func Error(v ...interface{}) {
	Default.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	Default.Errorf(format, v...)
}

func Warn(v ...interface{}) {
	Default.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	Default.Warnf(format, v...)
}

func Info(v ...interface{}) {
	Default.Info(v...)
}

func Infof(format string, v ...interface{}) {
	Default.Infof(format, v...)
}

func Debug(v ...interface{}) {
	Default.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	Default.Debugf(format, v...)
}

func Trace(v ...interface{}) {
	Default.Trace(v...)
}

func Tracef(format string, v ...interface{}) {
	Default.Tracef(format, v...)
}
