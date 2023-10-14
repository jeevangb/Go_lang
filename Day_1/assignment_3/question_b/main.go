package main

import (
	"GoProgrammes/Desktop/go_assignments/Day_1/assignment_3/question_b/temperature"
	"fmt"
)

func main() {
	var num float32
	fmt.Println("Enter temp in Fahrenheit")
	fmt.Scan(&num)
	res := temperature.Temp(num)
	fmt.Println(res)

}
