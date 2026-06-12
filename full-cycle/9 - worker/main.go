package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker:", workerId)
		fmt.Println("Msg:", res)
		time.Sleep(time.Second)
	}
}

func main() {
	msg := make(chan int)
	go worker(1, msg) // go routine
	go worker(2, msg) // go routine

	for i := 0; i < 10; i++ {
		msg <- i
	}
}
