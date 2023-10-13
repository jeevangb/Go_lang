package main

import "fmt"

func main() {
	numbers := []int{1, 3, 7, 2, 9, 4}
	fmt.Println(numbers)
	for i := 0; i < len(numbers); i++ {
		numbers[i] *= 2

	}
	fmt.Println(numbers)
}
