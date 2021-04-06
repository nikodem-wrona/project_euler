package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	rns := []rune(s)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func main() {
	fmt.Println("### START ###")

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j

			stringFromProduct := strconv.Itoa(product)

			partOne := stringFromProduct[0:3]
			partTwo := stringFromProduct[3:]

			reversedPartTwo := reverse(partTwo)

			if partOne == reversedPartTwo {
				fmt.Println("PALINDROME : ", partOne+partTwo)
				fmt.Println("PRODUCT : ", product)
			}
		}
	}
}
