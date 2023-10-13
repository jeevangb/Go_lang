package main

import "fmt"

func main() {
	numbers := []int{10, 200, 30, 60}

	num := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > num {
			num = numbers[i]
		}

	}
	fmt.Println("max element in a slice", num)
}
