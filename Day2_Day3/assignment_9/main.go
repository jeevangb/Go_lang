package main

import "fmt"

type Shape interface {
	Area()
}
type Circle struct {
	rad float32
}
type Rectangle struct {
	len     int
	breadth int
}

func (c Circle) Area() float32 {
	resC := 3.14 * c.rad * c.rad
	return resC

}
func (r Rectangle) Area() float32 {
	resR := 2 * r.len * r.breadth
	return float32(resR)

}

func main() {
	c1 := Circle{
		rad: 5.5,
	}
	r1 := Rectangle{
		len:     2,
		breadth: 4,
	}

	fmt.Println(c1.Area())
	fmt.Println(r1.Area())

}
