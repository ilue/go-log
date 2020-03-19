package log

import (
	"fmt"
	"strings"
)

type Level int32

const (
	PANIC Level = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
	_levelCount
)

var _levelTexts = [_levelCount]string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
}

func (l Level) String() string {
	if l >= 0 && l < _levelCount {
		return _levelTexts[l]
	}
	return fmt.Sprintf("Level(%d)", int(l))
}

func (l *Level) Set(s string) error {
	for i, text := range _levelTexts {
		if strings.EqualFold(s, text) {
			*l = Level(i)
			return nil
		}
	}
	return fmt.Errorf("invalid log level: %s", s)
}
