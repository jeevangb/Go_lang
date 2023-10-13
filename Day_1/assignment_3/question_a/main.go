package main

import (
	"GoProgrammes/Desktop/go_assignments/Day_1/assignment_3/question_a/calculator"
	"fmt"
)

func main() {
	a := calculator.Add(10, 20)
	s := calculator.Sub(30, 20)
	m := calculator.Mul(7, 7)
	d := calculator.Div(25, 5)

	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(d)
}
