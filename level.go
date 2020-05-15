package log

import (
	"errors"
	"strconv"
	"strings"
)

type Level int32

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	_levelCount
)

var _levelText = [_levelCount]string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
}

func (l *Level) Set(s string) error {
	for i := Level(0); i < _levelCount; i++ {
		if strings.EqualFold(_levelText[i], s) {
			*l = i
			return nil
		}
	}
	return errors.New("invalid level: " + s)
}

func (l Level) String() string {
	if l < _levelCount {
		return _levelText[l]
	}
	return "Level(" + strconv.Itoa(int(l)) + ")"
}
