package main

import (
	"errors"
	"fmt"
)

func main() {
	// Podemos ter int8, int16, int32 e int64
	var n1 int16 = 100 //Podemos fazer assim
	fmt.Println(n1)
	// Ou podemos fazer assim
	n2 := 50000
	fmt.Println(n2)
	// E até assim...
	var n3 int = 65421
	fmt.Println(n3)
	// Ainda temos uints, da mesma forma, das arquiteturas 8 ao 64
	var n4 uint64 = 4500
	fmt.Println(n4)

	// Sobre Alias
	// int32 = rune
	var n5 rune = 12345
	fmt.Println(n5)
	// uint8 = byte
	var n6 byte = 2
	fmt.Println(n6)

	// float32, float64
	var float_1 float32 = 123.45
	fmt.Println("float32", float_1)

	var float_2 float64 = 123.45
	fmt.Println("float64", float_2)

	// qualquer numero, seja int ou float, pode ser declarado como abaixo:
	qualquer_numero := 12
	fmt.Println(qualquer_numero)
	qualquer_numero1 := 12.5
	fmt.Println(qualquer_numero1)
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

	// Obs: ao usar aspas simples, o go identifica o numero identificador da tabela ASCII
	// Obs2: em aspas simples, nunca conseguiremos usar mais de 1 caractere.
	char := 'J'
	fmt.Println(char)
	
	// Ao não declarar nada em uma variável string é retornado nada, assim como int retorna o numeral zero
	// Booleano: não tem muito segredo. declara com true ou false. Valor default do bool é false
	var booleana_ bool = true
	fmt.Println(booleana_)
	// erro com seu valor zero = <nil>
	var erro error
	fmt.Println(erro)
	// ou...
	var erro2 error = errors.New("errors.New = Erro personalizavel para tratamento de erros. Pacote nativo do Golang.")
	fmt.Println(erro2)
}
