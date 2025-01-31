package main

import "fmt"

func inverterSinal(num int) int {
	return num * -1
}

func inverterSinalPonteiro(num *int) {
	*num = *num	* -1
}

func main()  {
	numero := 20
	numInvertido := inverterSinal(numero)

	fmt.Println(numInvertido)

	novoNumero := 40
	inverterSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}