package main

import "fmt"

func soma_n_valores(numero ...int) int {
	total := 0
	for _, n := range numero {
		total += n
	}
	return total
}

func main() {
	soma_dos_numeros := soma_n_valores(1, 2 , 5 , 5 ,8)
	fmt.Println(soma_dos_numeros)
}
