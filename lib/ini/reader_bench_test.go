package ini

import (
	"io/ioutil"
	"testing"
)

//
// 999f056 With bytes.Buffer in functions
// BenchmarkParse-2  300   17007586 ns/op  6361586 B/op  78712 allocs/op/
//
// 22dcd07 Move buffer to reader
// BenchmarkParse-2  500   19534400 ns/op  4656335 B/op  81163 allocs/op
//
// Refactor parser to use bytes.Reader
// BenchmarkParse-2  20000    72338 ns/op    35400 B/op    550 allocs/op
//
func BenchmarkParse(b *testing.B) {
	reader := NewReader()
	src, err := ioutil.ReadFile(testdataInputIni)
	if err != nil {
		b.Fatal(err)
	}

	for x := 0; x < b.N; x++ {
		_, _ = reader.Parse(src)
	}
}
