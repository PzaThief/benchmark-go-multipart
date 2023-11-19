package main

import (
	"testing"
)

const testFile100MB = "test-100MB"
const testFile10MB = "test-10MB"
const testFile1MB = "test-1MB"
const testFile100KB = "test-100KB"
const testFile10KB = "test-10KB"
const testFile1KB = "test-1KB"
const url = "http://localhost:8080"

func BenchmarkOsPipeSend100MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		osPipeSend(testFile100MB, url)
	}
}

func BenchmarkZeroCopyLibSend100MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeroCopyLibSend(testFile100MB, url)
	}
}

func BenchmarkIoPipeSend100MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioPipeSend(testFile100MB, url)
	}
}

func BenchmarkNaiveSend100MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		naiveSend(testFile100MB, url)
	}
}

func BenchmarkOsPipeSend10MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		osPipeSend(testFile10MB, url)
	}
}

func BenchmarkZeroCopyLibSend10MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeroCopyLibSend(testFile10MB, url)
	}
}

func BenchmarkIoPipeSend10MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioPipeSend(testFile10MB, url)
	}
}

func BenchmarkNaiveSend10MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		naiveSend(testFile10MB, url)
	}
}

func BenchmarkOsPipeSend1MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		osPipeSend(testFile1MB, url)
	}
}

func BenchmarkZeroCopyLibSend1MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeroCopyLibSend(testFile1MB, url)
	}
}

func BenchmarkIoPipeSend1MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioPipeSend(testFile1MB, url)
	}
}

func BenchmarkNaiveSend1MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		naiveSend(testFile1MB, url)
	}
}

func BenchmarkOsPipeSend100KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		osPipeSend(testFile100KB, url)
	}
}

func BenchmarkZeroCopyLibSend100KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeroCopyLibSend(testFile100KB, url)
	}
}

func BenchmarkIoPipeSend100KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioPipeSend(testFile100KB, url)
	}
}

func BenchmarkNaiveSend100KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		naiveSend(testFile100KB, url)
	}
}

func BenchmarkOsPipeSend10KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		osPipeSend(testFile10KB, url)
	}
}

func BenchmarkZeroCopyLibSend10KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeroCopyLibSend(testFile10KB, url)
	}
}

func BenchmarkIoPipeSend10KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioPipeSend(testFile10KB, url)
	}
}

func BenchmarkNaiveSend10KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		naiveSend(testFile10KB, url)
	}
}

func BenchmarkOsPipeSend1KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		osPipeSend(testFile1KB, url)
	}
}

func BenchmarkZeroCopyLibSend1KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeroCopyLibSend(testFile1KB, url)
	}
}

func BenchmarkIoPipeSend1KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioPipeSend(testFile1KB, url)
	}
}

func BenchmarkNaiveSend1KB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		naiveSend(testFile1KB, url)
	}
}
