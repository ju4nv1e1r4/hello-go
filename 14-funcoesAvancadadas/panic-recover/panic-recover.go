package main

import "fmt"

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada")
	}
}

func aluno_aprov(n1, n2, n3 float64) bool{
	defer recuperar()
	media := (n1 + n2 + n3) / 3
	
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("a média é exatamente 6")
}

func main()  {
	fmt.Println(aluno_aprov(10, 3, 5))
}