package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura float32
}

type estudante struct 	{
	pessoa
	curso string
	semestre uint8
}

func main() {
	fmt.Println("## O mais próximo que o golang chega de uma 'Herança' ##")

	p1 := pessoa{"Juan", "Vieira", 28, 1.70}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia de Software", 1}
	fmt.Println(e1)
}