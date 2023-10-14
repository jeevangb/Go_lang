package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	ch1 := make(chan int, 6)
	ch2 := make(chan int, 6)
	ch3 := make(chan int, 6)

	wg.Add(5)
	go sender1(ch1, wg)
	go sender2(ch2, wg)
	go sender3(ch3, wg)

	go reciver1(ch1, wg)
	go reciver2(ch2, wg)

	wg.Wait()
	// time.Sleep(300 * time.Second)
	fmt.Println("end of the main programme")

}

func sender1(ch1 chan int, Wg *sync.WaitGroup) {
	defer Wg.Done()
	for i := 1; i <= 6; i++ {

		ch1 <- i
	}

	close(ch1)

}

func reciver1(ch1 chan int, Wg *sync.WaitGroup) {
	defer Wg.Done()

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

func sender2(ch2 chan int, Wg *sync.WaitGroup) {
	defer Wg.Done()
	for i := 6; i <= 10; i++ {

		ch2 <- i
	}
	close(ch2)

}

func reciver2(ch2 chan int, Wg *sync.WaitGroup) {
	defer Wg.Done()

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

func sender3(ch3 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 6; i++ {

		ch3 <- i
	}

	close(ch3)

}
