package log

import (
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkLogger(b *testing.B) {
	logger := NewLogger(DEBUG, ioutil.Discard)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("abc", i)
	}
}

func BenchmarkStdLogger(b *testing.B) {
	logger := log.New(ioutil.Discard, "", log.LstdFlags|log.Lmicroseconds)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Print("abc", i)
	}
}
