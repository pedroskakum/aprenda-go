package main

import (
	"fmt"
)

func main() {
	hello := make(chan string)

	go func() {
		hello <- "Hello World"
	}()

	/*
		Esse caso interessante imprime default,
		pois o select nao tem nada para ler do canal hello,
		pois no momento que esse select é executado,
		a thread ainda nao escreveu nada no canal hello,
		caso voce coloque um sleep antes do select,
		o resultado seria diferente, pois a thread ja
		teria escrito algo no canal hello, e o select conseguiria ler
		TESTE: time.Sleep(time.Millisecond)
	*/

	// time.Sleep(time.Millisecond)

	select {
	case x := <-hello:
		fmt.Println("Recebendo do canal hello ", x)
	default:
		fmt.Println("default")
	}
}
