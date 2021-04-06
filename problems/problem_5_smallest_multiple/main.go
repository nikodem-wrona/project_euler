package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	startTime := time.Now()
	for {
		fmt.Println("I : ", i)
		var modulo []int

		for j := 1; j <= 20; j++ {
			result := i % j

			fmt.Println("RESULT : ", result)

			modulo = append(modulo, result)
		}

		sum := 0

		for _, number := range modulo {
			sum = sum + number
		}

		if sum == 0 {
			fmt.Println("SMALLEST MULTIPLE : ", i)
			endTime := time.Now()
			elapsed := endTime.Sub(startTime)

			fmt.Println("TASK TOOK : ", elapsed)
			return
		}

		i = i + 1
	}
}
