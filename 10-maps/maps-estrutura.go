package main

import (
	"fmt"
)

func main() {
	usuario := map[string]string {
		"Nome": "Juan",
		"Sobrenome": "Vieira",
	}

	fmt.Println(usuario)

	// map aninhado

	usuario2 := map[string]map[string]string {
		"nome": {
			"Primero Nome": "Juan",
			"Ultimo Nome": "Lucas",
		},
	}

	// adicionando novos dados (nova chave de map aninhado)
	usuario2["faixa etaria"] = map[string]string{
		"idade": "vinte e oito",
	}

	fmt.Println(usuario2)

	// printando valores individuais
	fmt.Println(usuario["Nome"]) // printa somente o nome
	fmt.Println(usuario2["nome"]["Ultimo Nome"]) // acessando maps aninhados
	fmt.Println(usuario2["faixa etaria"]) // print de dados adicionados
}