package main

import "fmt"

func main() {

	var ec int
	var oc int

	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			ec++
		} else {
			oc++
		}
	}
	fmt.Println("even count", ec)
	fmt.Println("odd count", oc)
}
