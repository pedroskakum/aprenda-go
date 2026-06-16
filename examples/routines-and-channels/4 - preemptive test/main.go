package main

import (
	"fmt"
	"runtime"
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
		Go organiza as thread de forma cooperativa,
		espera a thread terminar para comecar outra,
		exceto se ela demorar demais, quando isso acontece,
		Go se torna preemptivo

		Nesse caso, uma thread com um loop infinito de forma
		cooperativa travaria o programa, entao, a partir da versao 1.3,
		o modo preemptivo passa a valer para interromper execucoes muito longas
		e termina a execucao
	*/

	runtime.GOMAXPROCS(1)
	fmt.Println("Comecou")

	// funcao anonima e auto executavel rodando um loop infinito numa thread separada (travando)
	go func() {
		for {

		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Terminou")

}
