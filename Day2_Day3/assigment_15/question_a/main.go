package main

import (
	"fmt"
	"time"
)

// var Wg = &sync.WaitGroup{}

func main() {
	ch := make(chan int)
	go sender(ch)
	go reciver(ch)
	fmt.Println("end of the main programme")
	time.Sleep(3 * time.Second)

}

func sender(ch chan int) {
	for i := 1; i <= 6; i++ {

		ch <- i
	}

	// close(ch)

}

func reciver(ch chan int) {

	a := <-ch
	b := <-ch
	c := <-ch
	d := <-ch
	e := <-ch

	fmt.Println("data recived", a)
	fmt.Println("data recived", b)
	fmt.Println("data recived", c)
	fmt.Println("data recived", d)
	fmt.Println("data recived", e)

	// for v := range ch {
	// 	fmt.Println(v, <-ch)
	// }
}
