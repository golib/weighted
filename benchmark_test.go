package weighted

import (
	"math/rand/v2"
	"strconv"
	"testing"
)

func BenchmarkSW_Next(b *testing.B) {
	b.ReportAllocs()
	w := &SW{}
	for i := 0; i < 50; i++ {
		w.Add("item-"+strconv.Itoa(i), rand.IntN(100)+100)
	}

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w.Next()
		}
	})
}

func BenchmarkRRW_Next(b *testing.B) {
	b.ReportAllocs()
	w := &RRW{}
	for i := 0; i < 50; i++ {
		w.Add("item-"+strconv.Itoa(i), rand.IntN(100)+100)
	}

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w.Next()
		}
	})
}

func BenchmarkRandW_Next(b *testing.B) {
	b.ReportAllocs()
	w := NewRandW()
	for i := 0; i < 50; i++ {
		w.Add("item-"+strconv.Itoa(i), rand.IntN(100)+100)
	}

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w.Next()
		}
	})
}
