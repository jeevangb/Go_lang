package main

import "fmt"

type Calculate interface {
	area()
	perimeter()
}
type Rectangle struct {
	width  float32
	height float32
}

func (r Rectangle) area() {
	res := r.width * r.height
	fmt.Println("Area of rectangle:", res)
}
func (r Rectangle) perimeter() {
	res := 2 * (r.width + r.height)
	fmt.Println("Perimeter of rectangle:", res)
}

func main() {

	r1 := Rectangle{
		width:  10,
		height: 20,
	}
	r1.area()
	r1.perimeter()

}
