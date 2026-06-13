package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}
func main() {
	/*
		sem o "go" Eles estao rodando de forma serial,
		um esperando o outro, da forma tradicional

		com o "go" eles rodar de forma assyncrona,
		e nao e possivel ver o resultado sem sleep,
		para aguardar a thread
	*/
	go contador("a")
	go contador("b")
	time.Sleep(time.Second * 10)
}
