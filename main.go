package main

import (
	"fmt"
	"time"
)

func countToTen(c chan<- int) {
	for i := range [10]int{} {
		c <- i
		time.Sleep(1 * time.Second)
	}
	close(c)
}

func receive(c <-chan int) {
	for {
		a, ok := <-c
		if !ok {
			fmt.Println("we are done")
			break
		}
		fmt.Printf("received %d\n", a)
	}
}

func main() {
	// defer db.Close()
	// cli.Start()
	c := make(chan int)
	go countToTen(c)
	receive(c)

}
