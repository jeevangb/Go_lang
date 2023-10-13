package main

import "fmt"

func main() {
	fmt.Println("enter the number")
	var num int
	fmt.Scan(&num)

	fmt.Println(checkEvenOrOdd(num))
}
func checkEvenOrOdd(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
