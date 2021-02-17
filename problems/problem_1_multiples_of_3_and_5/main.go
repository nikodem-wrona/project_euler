// TASK:
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

type input []int64

func main() {
	inputSlice := generateInput(999)
	multiplesOf3or5 := inputSlice.getMultiplesOf3or5()
	sumOfMultiplesOf3or5 := multiplesOf3or5.sumAllNumbers()

	fmt.Println("Multiples of 3 or 5 :", multiplesOf3or5)
	fmt.Println("Sum of multiples of 3 or 5 : ", sumOfMultiplesOf3or5)
}

func (numbers input) getMultiplesOf3or5() input {
	var multiples []int64

	for _, number := range numbers {
		if number%3 == 0 || number%5 == 0 {
			multiples = append(multiples, number)
		}
	}

	return multiples
}

func (numbers input) sumAllNumbers() int64 {
	var sum int64 = 0

	for _, number := range numbers {
		sum = sum + number
	}

	return sum
}

func generateInput(amountOfNumbersInSlice int64) input {
	var numbers input

	for i := 0; i <= int(amountOfNumbersInSlice); i++ {
		numbers = append(numbers, int64(i))
	}

	return numbers
}
