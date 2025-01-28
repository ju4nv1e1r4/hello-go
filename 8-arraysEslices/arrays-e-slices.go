package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("## Arrays e Slices ##")

	var array1 [5]int // declarando array
	// OBS: arrays em go, só podem ter um tipo
	
	array1[0] = 5 // populando uma array
	// se você não popular as outras posições, por default, como tudo no go, valor-zero
	fmt.Println(array1)

	// outra forma de declarar array...
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5" }
	fmt.Println(array2)

	// declarar uma array tornando-a mais "flexível" por assim dizer
	array3 := [...]int{1, 2, 6, 5, 8, 10} // colocando os 3 pontos, o tamanho da array fica medido pela
	fmt.Println(array3) 				  // quantidade de elementos populados

	// Slices - são uma estrutura masi flexível de fato
	// flexível por tamanho, mas não por tipagem
	slice1 := []int{1, 2, 516, 55, 5}
	fmt.Println(slice1)

	// como ver o tipo?

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	slice1 = append(slice1, 18186)
	fmt.Println(slice1)

	// OBS² = Slice é uma fatia de um array
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "posição alterada"
	fmt.Println(slice2)

}