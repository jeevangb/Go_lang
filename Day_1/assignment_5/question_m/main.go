package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4, 9, 9}
	s2 := []int{1, 2, 3, 4, 5}
	if len(s1) != len(s2) {
		fmt.Println("slices are not equal")
		return
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			fmt.Println("false")
			return

		}

	}
	fmt.Println("true")
}
