package main

import (
	"GoProgrammes/Desktop/go_assignments/Day2_Day3/assignment_8/question_c/model"
	"GoProgrammes/Desktop/go_assignments/Day2_Day3/assignment_8/question_c/person"
)

func main() {
	person1 := model.Person{
		Age:  23,
		Name: "jeevan",
	}
	person2 := model.Person{
		Age:  34,
		Name: "afthab",
	}

	person.PrintPerson(person1)
	person.PrintPerson(person2)

}
