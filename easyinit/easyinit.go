package easyinit

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/ilue/go-log"
)

var (
	_filename   string
	_maxSize    int
	_maxAge     int
	_maxBackups int
	_noConsole  bool
)

func init() {
	flag.StringVar(&_filename,
		"log-filename",
		filepath.Join(os.TempDir(), filepath.Base(os.Args[0])+".log"),
		"File to write logs to")

	flag.IntVar(&_maxSize,
		"log-maxSize",
		0,
		"Maximum size in megabytes of the log file before it gets rotated")

	flag.IntVar(&_maxAge,
		"log-maxAge",
		0,
		"Maximum number of days to retain old log files based on the timestamp encoded in their filename")

	flag.IntVar(&_maxBackups,
		"log-maxBackups",
		0,
		"Maximum number of old log files to retain")

	flag.BoolVar(&_noConsole,
		"log-noConsole",
		false,
		"Don't write logs to console")

	log.Default = log.Output(&bootstrap{})
}

type bootstrap struct {
	mu  sync.Mutex
	out io.Writer
}

func (b *bootstrap) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.out == nil {
		var outs []io.Writer
		if _filename != "" {
			outs = append(outs, &lumberjack.Logger{
				Filename:   _filename,
				MaxSize:    _maxSize,
				MaxAge:     _maxAge,
				MaxBackups: _maxBackups,
			})
		}
		if !_noConsole {
			outs = append(outs, log.SyncWriter(os.Stderr))
		}
		if len(outs) > 0 {
			b.out = io.MultiWriter(outs...)
		} else {
			b.out = ioutil.Discard
		}
		log.Default = log.WithOutput(b.out)
	}

	return b.out.Write(p)
}
