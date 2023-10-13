package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 9, 7, 4}
	fmt.Println("enter the number")
	var num int

	fmt.Scan(&num)

	res := check(num, numbers)
	fmt.Println(res)
}
func check(num int, a []int) int {
	for i := 0; i < len(a); i++ {

		if num == a[i] {
			return i
		}

	}
	return -1
}
