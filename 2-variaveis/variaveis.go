package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "Essa variável foi declarada explicítamente uma string"
	variavel2 := "Essa variável foi declarada uma string, implícitamente"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Empacotando as ..."
		variavel4 string = "... variáveis explícitamente."
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Agora estou empacotando...", "...as variáveis implícitamente."
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Esta é a minha primeira constante. Será que é equivalente a tuplas no Python?"
	fmt.Println(constante1)
}