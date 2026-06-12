package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int)

	/*
		O loop nesse caso so vai passar para a proxima
		interacao depois que alguem ler o valor do canal queue,
		pois o canal queue é bloqueante, ou seja, ele bloqueia a thread
	*/
	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			queue <- i
			i++
		}
	}()

	for x := range queue {
		fmt.Println(x)
	}

}
