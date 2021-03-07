package main

import (
	"fmt"
)

// sieveOfEratosthenes returns list of primes less than n
func sieveOfEratosthenes(n int) []int {
	b := make([]bool, n)
	var primes []int

	for i := 2; i < n; i++ {
		if b[i] == true {
			continue
		}

		primes = append(primes, i)

		for k := i * i; k < n; k += i {
			b[k] = true
		}
	}

	return primes
}

func primeFactorization(number int, primes []int) []int {
	var factors []int
	var dividingResult int = number

	for _, p := range primes {
		for (dividingResult % p) == 0 {

			dividingResult = dividingResult / p
			factors = append(factors, p)
		}
	}

	return factors
}

func main() {
	number := 600851475143
	primes := sieveOfEratosthenes(10000)
	factors := primeFactorization(number, primes)
	largestPrimeFactor := factors[len(factors)-1]
	fmt.Printf("Largest prime factor of number : %d is : %d", number, largestPrimeFactor)
}
