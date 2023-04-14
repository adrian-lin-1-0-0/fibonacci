package main

import (
	"testing"
)

const n = 10000

func BenchmarkRecursionTree(b *testing.B) {
	recursionTree(n)
}

func BenchmarkFastDoubling(b *testing.B) {
	fastDoubling(n)
}

func BenchmarkRecursionChannel(b *testing.B) {
	recursionChannel(n)
}

func BenchmarkTailCall(b *testing.B) {
	tailCall(n)
}

func BenchmarkForLoop(b *testing.B) {
	forLoop(n)
}
