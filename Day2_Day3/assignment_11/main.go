package main

import "fmt"

type Address struct {
	street  string
	city    string
	zipcode int
}
type Person struct {
	name string
	a    Address
}

func main() {
	p1 := Person{
		name: "Jeevan",
		a: Address{
			street:  "Bairasandra",
			city:    "Bnaglore",
			zipcode: 5600001,
		},
	}
	fmt.Println(p1)

}
