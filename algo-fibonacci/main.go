package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func fibonacci(n uint) uint {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func fib(n int, memo map[int]int) int {
	if val, ok := memo[n]; ok {
		return val
	}

	if n <= 1 {
		return n
	}

	memo[n] = fib(n-1, memo) + fib(n-2, memo)

	return memo[n]
}

func main() {
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	fibonacci(25)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	start = time.Now()
	r = new(big.Int)
	memo := make(map[int]int)
	fib(25, memo)

	elapsed = time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
