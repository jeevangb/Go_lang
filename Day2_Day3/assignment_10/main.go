package main

import "fmt"

func main() {

	res := sum(2, 3, 7)
	fmt.Println(res)

}
func sum(nums ...int) int { //variadic paramter
	sum := 0
	for _, v := range nums {
		sum += v

	}
	return sum

}
