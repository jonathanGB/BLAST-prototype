package main

import "testing"

func benchmarkKmerX(k int, q string, b *testing.B) {
	db := initDB()
	data := &blastJSON{
		q,
		false,
		k,
	}

	for n := 0; n < b.N; n++ {
		data.getBlastOutput(db)
	}
}

func BenchmarkKmer3_20(b *testing.B) {
	benchmarkKmerX(3, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer4_20(b *testing.B) {
	benchmarkKmerX(4, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer5_20(b *testing.B) {
	benchmarkKmerX(5, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer6_20(b *testing.B) {
	benchmarkKmerX(6, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer7_20(b *testing.B) {
	benchmarkKmerX(7, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer8_20(b *testing.B) {
	benchmarkKmerX(8, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer9_20(b *testing.B) {
	benchmarkKmerX(9, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer10_20(b *testing.B) {
	benchmarkKmerX(10, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer11_20(b *testing.B) {
	benchmarkKmerX(11, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer12_20(b *testing.B) {
	benchmarkKmerX(12, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer13_20(b *testing.B) {
	benchmarkKmerX(13, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer14_20(b *testing.B) {
	benchmarkKmerX(14, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer15_20(b *testing.B) {
	benchmarkKmerX(15, "AGCGAATGAGTAAACTGTAC", b)
}

func BenchmarkKmer3_40(b *testing.B) {
	benchmarkKmerX(3, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer4_40(b *testing.B) {
	benchmarkKmerX(4, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer5_40(b *testing.B) {
	benchmarkKmerX(5, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer6_40(b *testing.B) {
	benchmarkKmerX(6, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer7_40(b *testing.B) {
	benchmarkKmerX(7, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer8_40(b *testing.B) {
	benchmarkKmerX(8, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer9_40(b *testing.B) {
	benchmarkKmerX(9, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer10_40(b *testing.B) {
	benchmarkKmerX(10, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer11_40(b *testing.B) {
	benchmarkKmerX(11, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer12_40(b *testing.B) {
	benchmarkKmerX(12, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer13_40(b *testing.B) {
	benchmarkKmerX(13, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer14_40(b *testing.B) {
	benchmarkKmerX(14, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}

func BenchmarkKmer15_40(b *testing.B) {
	benchmarkKmerX(15, "TGGACCAAATCGGCCTTTCTGTAGGGGACCGGGTCTTAGG", b)
}
