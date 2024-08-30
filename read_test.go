package goreadinput_test

import (
	"testing"

	goreadinput "github.com/nailanmnabil/go-read-input"
)

func BenchmarkReadAllSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goreadinput.ReadAll("data_sm.csv")
	}
}

func BenchmarkBufferedReadingSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goreadinput.BufferedReading("data_sm.csv")
	}
}

func BenchmarkMmapSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goreadinput.Mmap("data_sm.csv")
	}
}

func BenchmarkReadAllMed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goreadinput.ReadAll("data_md.csv")
	}
}

func BenchmarkBufferedReadingMed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goreadinput.BufferedReading("data_md.csv")
	}
}

func BenchmarkMmapMed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goreadinput.Mmap("data_md.csv")
	}
}

func BenchmarkReadAllBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goreadinput.ReadAll("data_big.csv")
	}
}

func BenchmarkBufferedReadingBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goreadinput.BufferedReading("data_big.csv")
	}
}

func BenchmarkMmapBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goreadinput.Mmap("data_big.csv")
	}
}
