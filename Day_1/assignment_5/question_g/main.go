package main

import "fmt"

func main() {
	number := []int{1, 6, 8, 9, 7}
	f := 0
	l := len(number) - 1
	var temp int
	fmt.Println("before reverse")
	fmt.Println(number)
	for f < l {
		temp = number[f]
		number[f] = number[l]
		number[l] = temp
		f++
		l--

	}
	fmt.Println("after reverse")
	fmt.Println(number)

}
