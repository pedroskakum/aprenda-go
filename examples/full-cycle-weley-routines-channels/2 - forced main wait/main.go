package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
	}
}
func main() {
	contador("sem go routine")
	/*
		Se colocarmos o time.Sleep no final de 1 segundo,
		os prints da thread "go" vao rodar, se remover o sleep
		eles vao deixar de ser executados novamente
	*/
	go contador("com go routine")
	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
	time.Sleep(time.Second)
}
