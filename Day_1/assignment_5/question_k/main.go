package main

import "fmt"

func main() {

	numbers := []int{1, 2, 2, 3, 1}
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				numbers = deleteElement(numbers, j)
			}

		}

	}
	fmt.Println(numbers)

}
func deleteElement(rcvslice []int, j int) []int {
	rcvslice = append(rcvslice[:j], rcvslice[j+1:]...)
	return rcvslice
}
