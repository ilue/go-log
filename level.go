package log

import (
	"fmt"
)

type Level int32

const (
	PANIC Level = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	_levelCount
)

var _levelTexts = []string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
}

func (l Level) String() string {
	if l >= 0 && l < _levelCount {
		return _levelTexts[l]
	}
	return fmt.Sprintf("Level(%d)", int(l))
}
