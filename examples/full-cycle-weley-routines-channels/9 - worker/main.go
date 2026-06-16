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

	/*
		quanto mais workers eu criar, mais rapido o processamento da fila vai acontecer,
		pois cada worker vai processar uma parte da fila, de forma independente,
		e isso e uma das vantagens de usar go routines, pois elas permitem que voce crie
		muitas threads de forma facil, e isso e muito util para processar filas de tarefas
	*/
	go worker(1, msg) // go routine
	go worker(2, msg) // go routine

	for i := 0; i < 10; i++ {
		msg <- i
	}
}
