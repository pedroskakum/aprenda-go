package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá mundo!")
}

func main() {
	http.HandleFunc("/", home)

	fmt.Println("Servidor rodando em :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
