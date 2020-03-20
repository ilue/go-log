package log

import (
	"io"
	"sync"
)

type syncWriter struct {
	mu sync.Mutex
	w  io.Writer
}

func SyncWriter(w io.Writer) io.Writer {
	return &syncWriter{w: w}
}

func (s *syncWriter) Write(p []byte) (int, err) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.w.Write(p)
}
