package main

import "fmt"

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

func main() {
	primes := sieveOfEratosthenes(2000000)
	var sum int64 = 0;

	for _, prime := range primes {
		sum += int64(prime);
	}

	fmt.Println("Sum : ", sum)
}
