package main

import "fmt"

type user struct {
	nome  string
	idade uint
}

func (u user) salvar() {
	fmt.Printf("simulação: Salvando os dados do usuário %s no database\n", u.nome)
}

func main()  {
	usuario1 := user{"Juan", 28}
	usuario1.salvar()
}