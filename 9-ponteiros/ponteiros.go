package main

import "fmt"

func main () {
	fmt.Println("## Aula sobre Ponteiros ##")

	// Um ponteiro é uma referência de memória
	// Exemplo:

	var var1 int = 10
	var var2 int = var1
	var1++
	fmt.Println(var1, var2)
	// Ao imprimirmos isso na tela do terminal, 
	// podemos ver que var1 foi somado, mas apesar
	// de var2 ser = var1, não é verdade que são a mesma coisa.
	// Na verdade, var2 é uma copia de var1 mas em outro
	// endereço de memória. Como se fossem clones.
	// Se um muda, o outro não muda.

	var var3 int        // isso é um integer
	var ponteiro1 *int  // isso é um ponteiro de um integer
	
	var3 = 100 
	ponteiro1 = &var3 // pra isso aqui dar certo usasse um & comercial antes do atributo do ponteiro

	// acima temos duas coisas diferentes
	fmt.Println(ponteiro1) // aqui será impresso o endereço de memória
	fmt.Println(*ponteiro1) // caso eu queira ver o valor do ponteiro, devo usar * antes, isso é chamado desreferenciação
	// Desreferenciar é uma operação que permite a um programador acessar o valor armazenado em um endereço de memória
	// indicado por um ponteiro. 



}