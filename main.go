package main

import (
	"fmt"
	"time"
)

func countToTen(c chan int) {
	for i := range [10]int{} {
		c <- i
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// defer db.Close()
	// cli.Start()
	c := make(chan int)
	go countToTen(c)
	for{
		
		a := <-c
		fmt.Printf("received %d\n", a)
	}
}
