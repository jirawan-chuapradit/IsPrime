package main

import "testing"

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(9897)
	}
}

func BenchmarkIsPrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime2(9897)
	}
}
