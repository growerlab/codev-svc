package main

import (
	"testing"
)

func Benchmark_gogit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gogit()
	}
}
func Benchmark_git2go(b *testing.B) {
	for i := 0; i < b.N; i++ {
		git2go()
	}
}
