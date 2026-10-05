package main

import (
	"fmt"
)

func main() {
	sum := sumup(1, 10, 15)

	numbers := []int{1, 10, 15}
	anotherSum := sumup(0, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum + startingValue
}
