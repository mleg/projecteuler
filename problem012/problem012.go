package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	start := time.Now()
	n := 1 // n goes for triangle numbers
	primes := []int{2} // for accumulating prime numbers

	for i := 2; ; i++ {
		n += i
		if divisors(n, &primes) > 500 {
			fmt.Printf("The answer is %v, %v\n", n, time.Since(start))
			return
		}
	}
}

func divisors(n int, primes *[]int) int {

	primeFactors := primeFactors(n , primes)
	result := 1
	for _, power := range primeFactors {
		result *= (power + 1)
	}
	return result
}

func primeFactors(n int, primes *[]int) map[int]int {

	result := make(map[int]int)

	OuterLoop:
	for n != 1 {
		limit := sqrt(n)
		for _, divisor := range *primes {
			if divisor > limit {
				break OuterLoop
			} else if n%divisor == 0 {
				result[divisor]++
				n = n / divisor
				continue OuterLoop
			}
		}
		// not enough primes yet
		*primes = append(*primes, nextPrime(*primes))
	}

	if n != 1 {
		result[n]++
	}

	return result
}

func sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func nextPrime(primes []int) int {

	n := primes[len(primes)-1] + 1
	if n%2 == 0 { // only odd numbers are eligible
		n++
	}

	// Finding n as the first number not divisable by any of previous primes
	for ; ; n += 2 {
		limit := sqrt(n)
		for _, divisor := range primes {
			if divisor > limit {
				return n // so n is a prime number
			} else if n%divisor == 0 {
				break // go to next n
			}
		}
	}
}
