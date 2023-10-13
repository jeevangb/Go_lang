package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter number which you want to search")
	fmt.Scan(&num)

	numbers := []int{10, 36, 24, 376, 78}
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == num {
			fmt.Println("number found")
			return

		}

	}
	fmt.Println("number not found")

}
