package main

import "fmt"

func fibonacci(posicao uint) uint{
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao -2) + fibonacci(posicao-1)
}


func main() {
	position := uint(8)
	fmt.Println(fibonacci(position))
}