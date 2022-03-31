package main

import "testing"

const filename = hundredK

func BenchmarkTopK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getRanking(10, filename)
	}
}

func BenchmarkTopKOptimised(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getRankingOptimized(10, filename)
	}
}
