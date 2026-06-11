// crie um array de 5 posicioes, preencha-o e demonstre
package main

import "fmt"

func main() {
	meuarray := [5]int{1, 2, 3, 4, 5}

	// fmt.Println(meuarray)
	for _, v := range meuarray {
		fmt.Printf("%T\n", v)
		fmt.Println(v, "\n")
	}
}
