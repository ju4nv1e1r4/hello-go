package main

import "fmt"

func somar(n1 int64, n2 int64) int64 {
	return n1 + n2
}

func Calculos(n1, n2 int32) (int32, int64) {
	soma := n1 + n2
	subt := n1 - n2
	return soma, int64(subt)
}

func main() {
	soma := somar(100, 200)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("usando função aninhada")

	results_soma, results_subt := Calculos(10, 20)
	fmt.Println(results_soma, results_subt)
}
