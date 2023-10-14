package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 6)
	ch2 := make(chan int, 6)
	ch3 := make(chan int, 6)

	go sender1(ch1)
	go sender2(ch2)
	go sender3(ch3)

	go reciver1(ch1)
	go reciver2(ch2)

	time.Sleep(300 * time.Second)
	fmt.Println("end of the main programme")

}

func sender1(ch1 chan int) {
	for i := 1; i <= 6; i++ {

		ch1 <- i
	}

	close(ch1)

}

func reciver1(ch1 chan int) {

	a := <-ch1
	b := <-ch1
	c := <-ch1
	d := <-ch1
	e := <-ch1

	fmt.Println("data recived", a)
	fmt.Println("data recived", b)
	fmt.Println("data recived", c)
	fmt.Println("data recived", d)
	fmt.Println("data recived", e)
}

func sender2(ch2 chan int) {
	for i := 1; i <= 6; i++ {

		ch2 <- i
	}
	close(ch2)

}

func reciver2(ch2 chan int) {

	a := <-ch2
	b := <-ch2
	c := <-ch2
	d := <-ch2
	e := <-ch2

	fmt.Println("data recived", a)
	fmt.Println("data recived", b)
	fmt.Println("data recived", c)
	fmt.Println("data recived", d)
	fmt.Println("data recived", e)
}

func sender3(ch3 chan int) {
	for i := 1; i <= 6; i++ {

		ch3 <- i
	}

	close(ch3)

}
