package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, substracao int) {
	soma = n1 + n2
	substracao = n1 - n2
	return
}

func main() {
	soma, substracao := calculosMatematicos(10, 20)
	fmt.Println(soma, substracao)
}
