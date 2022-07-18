package main

import "testing"

func BenchmarkForLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForLoop([]string{"Hi", "there", "user", "nice", "to", "see", "you"})
	}
}

func BenchmarkRangeLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeLoop([]string{"Hi", "there", "user", "nice", "to", "see", "you"})
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"Hi", "there", "user", "nice", "to", "see", "you"})
	}
}
