package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	rua string
	numero uint8
}

func main() {
	fmt.Println("## Aprendendo sobre Structs ##")

	// primeira forma de popular uma struct
	fmt.Println(">> Primeira forma de popular uma struct")
	var u usuario
	u.nome = "Juan Vieira"
	u.idade = 28
	fmt.Println(u) // se imprimir sem popular a struct, terá um valor zero. saida terminal >> { 0}

	enderecoUsuario := endereco{"Rua brasil", 101}

	// segunda forma de popular uma struct
	fmt.Println(">> Segunda forma de popular uma struct")
	u2 := usuario{"Juan Vieira", 29, enderecoUsuario}
	fmt.Println(u2)

	// terceira forma de popular uma struct
	fmt.Println(">>Terceira forma de popular uma struct")
	u3 := usuario{nome: "juan"}
	fmt.Println(u3)
}