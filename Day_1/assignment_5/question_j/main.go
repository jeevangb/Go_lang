package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{8, 2, 4}

	fmt.Println(s1)
	fmt.Println(s2)
	s1 = append(s1, s2...)
	fmt.Println("after append")
	fmt.Println(s1)

}
