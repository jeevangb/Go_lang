package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	s1 := Student{
		name: "jeevan",
		age:  22,
	}
	// s2 := Student{
	// 	name: "poorvi",
	// 	age:  25,
	// }
	// s3 := Student{
	// 	name: "afthab",
	// 	age:  23,
	// }

	update(&s1)
	fmt.Println(s1)

}

func update(s *Student) {
	s.age += 1
	fmt.Println(s.age)

}
