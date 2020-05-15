package log

import (
	"io/ioutil"
	"testing"
)

func BenchmarkLogger(b *testing.B) {
	l := New().Output(ioutil.Discard).Logger()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		l.Info("Some log text",
			"str1", "foo",
			"str2", "bar",
			"int1", 123,
			"int2", -456,
		)
	}
}
