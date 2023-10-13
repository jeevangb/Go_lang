package main

import "fmt"

func main() {

	type Employee struct {
		id   int
		name string
		age  int
		city string
	}

	e1 := Employee{1, "jeebvan", 23, "chitradurga"}
	e2 := Employee{2, "afthab", 24, "Kodogu"}
	e3 := Employee{
		id:   3,
		name: "poorvi",
		age:  22,
		city: "Banglore",
	}

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}
