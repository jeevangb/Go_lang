package main

import "fmt"

type Animal interface {
	MakeSound()
}

type Dog struct {
	sound string
}
type Cat struct {
	sound string
}

func (d Dog) MakeSound() {
	fmt.Println(d.sound)

}
func (c Cat) MakeSound() {
	fmt.Println(c.sound)

}

func main() {
	d := Dog{
		sound: "bou bou",
	}

	c := Cat{
		sound: "meow",
	}
	c.MakeSound()
	d.MakeSound()

}
