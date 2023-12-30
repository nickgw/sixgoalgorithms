package main

import "fmt"

func main() {
	var n int64
	for n = 0; n <= 21; n++ {
		fmt.Printf("%3d! = %20d\n", n, factorial(n))
	}
	fmt.Println()
}

func factorial(n int64) int64 {
	if n < 1 {
		return int64(1)
	} else {
		return n * factorial(int64(n-1))
	}
}
