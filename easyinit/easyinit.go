package easyinit

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilue/go-log"
	"github.com/natefinch/lumberjack"
)

var (
	_filename   string
	_maxSize    int
	_maxBackups int
	_noConsole  bool
)

func init() {
	flag.StringVar(&_filename,
		"log-filename",
		filepath.Join(os.TempDir(), filepath.Base(os.Args[0])+".log"),
		"")

	flag.IntVar(&_maxSize,
		"log-maxsize",
		0,
		"")

	flag.IntVar(&_maxBackups,
		"log-maxbackups",
		0,
		"")

	flag.BoolVar(&_noConsole,
		"log-noconsole",
		false,
		"")

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
		log.Default = log.Default.Output(b.out)
	}

	return b.out.Write(p)
}
