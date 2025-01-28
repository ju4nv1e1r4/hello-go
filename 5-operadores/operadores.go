package main

import "fmt"

func main() {
	// OPERADORES ARITMETICOS
	soma := 10 + 2
	subtracao := 12 - 2
	divisao := 12 / 2
	multiplicacao := 6 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	// OBS: NAÕ PODEMOS FAZER OPERAÇÕES COM TIPOS DIFERENTES. EXEMPLO: INT32 + INT64. GO É FORTEMENTE TIPADO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 > 2)
	fmt.Println(1 == 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 != 2)

	// OPERADORES LÓGICOS
	fmt.Println(false && true) // and
	fmt.Println(false || true) // or
	fmt.Println(!false) // not

	// OPERADORES UNÁRIOS
	numero := 10
	numero++ // incrementa + 1
	numero += 15 //incrementa quantos você quiser
	fmt.Println(numero)

	numero-- //subtrai - 1 
	numero -= 2 // subtrai quantos quiser
	fmt.Println(numero)

	numero *= 3 // multiplica por 3
	fmt.Println(numero)

	numero /= 2 // divide por 2
	fmt.Println(numero)

	numero %= 2
	fmt.Println(numero)
}
