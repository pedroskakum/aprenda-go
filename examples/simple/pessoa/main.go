package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	meumapa := make(map[string]pessoa)

	meumapa["Pimentão"] = pessoa{
		nome:      "Renata",
		sobrenome: "Pimentão",
		sabores:   []string{"Pistache", "Morango", "Baunilha"},
	}

	meumapa["da Prussia"] = pessoa{
		nome:      "Frederico",
		sobrenome: "da Prussia",
		sabores:   []string{"Sabao em po", "Pe de coelho", "Feijao"},
	}

	for _, v := range meumapa {
		fmt.Println("Meu nome e ", v.nome, v.sobrenome, "e meu sabores favoritos sao")

		for _, m := range v.sabores {
			fmt.Println("-", m)
		}
		fmt.Println()
	}
}
