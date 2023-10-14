package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 4}

	var sum int
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			sum = sum + numbers[i]
		}

	}
	fmt.Println("even sum :", sum)

}
