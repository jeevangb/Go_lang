package main

import "fmt"

func main() {
	numbers := []int{}

	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	//  numbers:=appen(numbers,5)\
	numbers = append(numbers, 5)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers, 10)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers, 15)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers, 20)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers, 25)
	fmt.Println(numbers)

	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println(numbers)

}
