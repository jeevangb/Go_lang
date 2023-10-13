package main

import "fmt"

func main() {

	var num int
	fmt.Println("enter the number")
	fmt.Scan(&num)
	res := factorial(num)
	fmt.Println(res)

}

var p int = 1

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)

}
