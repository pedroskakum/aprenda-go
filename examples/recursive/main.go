package main

import "fmt"

func factorial(numero int) int {
	if numero == 1 {
		return numero
	}
	var temp = numero
	temp--
	return numero * factorial(temp)
}

func main() {
	var numero int
	fmt.Scan(&numero)
	fmt.Println(factorial(numero))
}
