package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	// Forma 1:
	for i < 10 {
		i++
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	
	// Forma 2:
	fmt.Println(i)

	for j:= 0; j < 10; j+=2 {
		println(j)
		time.Sleep(time.Second)
	}

	// Forma 3(com range percorrendo array (ou slice se preferir)):
	pessoas := [3]string{"Juan", "Luana", "Lucas"}

	for _, nome := range pessoas {
		fmt.Println(nome)
	}

	// Forma 3: percorrendo letras de uma string 
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
	
	// Forma 3: percorrendo maps
	usuario := map[string]string {
		"Nome": "Juan",
		"Sobrenome": "Vieira",
	}

	for k, v := range usuario {
		fmt.Println(k,":", v)
	}

	// NÃO É POSSÍVEL FAZER RANGE EM STRUCTS
}