package log

import (
	"io"
)

type LoggerBuilder struct {
	logger Logger
}

func New() *LoggerBuilder {
	return &LoggerBuilder{
		logger: defaultLogger(),
	}
}

func (b *LoggerBuilder) Level(v Level) *LoggerBuilder {
	b.logger.level = v
	return b
}

func (b *LoggerBuilder) Output(w io.Writer) *LoggerBuilder {
	b.logger.out = w
	return b
}

func (b *LoggerBuilder) Fields(args ...interface{}) *LoggerBuilder {
	b.logger.context = appendFields(b.logger.context, args)
	return b
}

func (b *LoggerBuilder) Logger() Logger {
	return b.logger
}
