package main

import "fmt"

func funClosure() func() {
	texto := "Este é o texto da função Closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return  funcao
}

func main()  {
	texto := "Este é o texto  da função main"
	fmt.Println(texto)

	novaFuncao := funClosure()
	novaFuncao()
}