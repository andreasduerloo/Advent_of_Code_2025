package day03

import "testing"

var blackhole int

func BenchmarkOptimizedHighestJoltage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := highestJoltage([]int{3, 7, 3, 3, 4, 4, 4, 4, 4, 4, 3, 3, 7, 2, 4, 4, 3, 4, 1, 4, 6, 3, 4, 5, 2, 4, 6, 3, 6, 4, 4, 2, 3, 4, 4, 9, 3, 3, 5, 4, 1, 4, 4, 5, 8, 4, 4, 3, 3, 3, 4, 4, 4, 2, 5, 4, 4, 4, 4, 5, 3, 5, 3, 4, 4, 5, 4, 4, 4, 4, 3, 4, 3, 4, 4, 4, 3, 2, 4, 3, 3, 5, 4, 5, 4, 4, 4, 6, 4, 2, 3, 3, 4, 3, 4, 4, 4, 4, 7, 2}, 0, 12)
		blackhole = result
	}
}

func BenchmarkUnoptimizedHighestJoltage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := highestJoltageFullRange([]int{3, 7, 3, 3, 4, 4, 4, 4, 4, 4, 3, 3, 7, 2, 4, 4, 3, 4, 1, 4, 6, 3, 4, 5, 2, 4, 6, 3, 6, 4, 4, 2, 3, 4, 4, 9, 3, 3, 5, 4, 1, 4, 4, 5, 8, 4, 4, 3, 3, 3, 4, 4, 4, 2, 5, 4, 4, 4, 4, 5, 3, 5, 3, 4, 4, 5, 4, 4, 4, 4, 3, 4, 3, 4, 4, 4, 3, 2, 4, 3, 3, 5, 4, 5, 4, 4, 4, 6, 4, 2, 3, 3, 4, 3, 4, 4, 4, 4, 7, 2}, 0, 12)
		blackhole = result
	}
}
