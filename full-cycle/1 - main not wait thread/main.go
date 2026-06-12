package main

import (
	"fmt"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
	}
}
func main() {
	contador("sem go routine")
	/*
		Cade sem print "com go routine"?
		quando eu coloquei "go" na frente do metodo contador,
		foi criado uma thread separada da funcao principal
		e como o programa executou de forma muito rapida
		nao deu tempo executar a thread antes do
		programa principal finalizar
	*/
	go contador("com go routine")
	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
}
