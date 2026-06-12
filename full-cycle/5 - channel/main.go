package main

import "fmt"

// THREAD1
func main() {
	/*
		chan = channel = forma de compatilhar valores entre threads diferentes
		THREAD1 <=> THREAD2
	*/
	hello := make(chan string)

	// THREAD2
	go func() {
		hello <- "Hello World"
	}()

	result := <-hello
	fmt.Println(result)
}
