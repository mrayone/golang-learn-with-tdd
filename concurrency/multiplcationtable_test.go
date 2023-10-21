package concurrency_test

import (
	"github.com/mrayone/learn-go/concurrency"
	"testing"
)

// BenchmarkPrintTable-16                 1        1204754073 ns/op
func BenchmarkPrintTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrency.PrintTable()
	}
}

// BenchmarkPrintMain-16                  2         607598010 ns/op is more fast
func BenchmarkPrintMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrency.PrintMain()
	}
}
