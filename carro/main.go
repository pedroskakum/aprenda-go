/*
Crie um novo tipo veiculo
	O tipo subjacente deve ser struct
	deve conter os campos: portas, cor
Crie dois novos tipos: caminhonete e sedan
	Os tipos subjacentes devem ser struct
	Ambos devem conter veiculo como struct embutido
	o tipo caminhonete deve conter um campo bool chamado tracaoNasQuatro
	o tipo sedan deve conter um campo bool chamado modeloLu
Usando os structs veiculo, caminhonete e sedan
	usando composite literal, crie um valor de tipo caminhonete e sedan e de valores  a seus campos
	demostre os valores
*/

package main

import "fmt"

type Veiculo struct {
	portas int
	cor    string
}

type Caminhonete struct {
	Veiculo
	tracaoNasQuatro bool
}

type Sedan struct {
	Veiculo
	modeloLuxo bool
}

func main() {
	Fusca := Veiculo{
		portas: 2,
		cor:    "branca",
	}

	// Declaracao posicional
	Corolla := Sedan{
		Veiculo{4, "prata"},
		false,
	}

	// Declaracao nomeada
	Hilux := Caminhonete{
		Veiculo: Veiculo{
			portas: 2,
			cor:    "preta",
		},
		tracaoNasQuatro: true,
	}

	fmt.Println(Fusca)
	fmt.Println(Corolla)
	fmt.Println(Hilux)

	// Voce pode usar o Caminhonete.Veiculo.cor diretamente como Caminhonete.cor
	fmt.Println(Hilux.cor)

}
