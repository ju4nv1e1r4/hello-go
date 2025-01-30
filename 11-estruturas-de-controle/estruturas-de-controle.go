package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle (famoso if/else)")

	valor := 5

	if valor > 15 {
		fmt.Println("É maior do que 15.")
	} else {
		fmt.Println("É menor ou igual a 15")
	}

	if outroValor := valor; outroValor > 5 {
		fmt.Println("É maior do que 5.")
	} else if outroValor < 5 {
		fmt.Println("É menor do que 5")
	} else {
		fmt.Println("É igual a 5")
	}
}