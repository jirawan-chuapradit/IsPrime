package main

import "testing"

func TestIsPrime(t *testing.T) {
	if !isPrime2(13) {
		t.Error("13 is not prime")
	}
}
