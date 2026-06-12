package main

import "fmt"

func main() {
	/*
		foverer esta aguardando para sempre,
		independentemente do resultado da thread
	*/
	foverer := make(chan string)

	go func() {
		x := true
		for {
			if x == true {
				// fmt.Println("x == true")
				continue
			}
		}
	}()

	fmt.Println("Aguardando para sempre")
	<-foverer
}
