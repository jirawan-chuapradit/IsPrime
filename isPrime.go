package main

import (
	"fmt"
	"math"
)

func main() {
	ans := isPrime(13)
	fmt.Println(ans)
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func isPrime2(x int) bool {
	for i := 2; i < int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
