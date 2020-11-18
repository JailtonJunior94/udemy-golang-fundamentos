package main

import "fmt"

func main() {
	/* Aritméticos */
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	/* Atribuição */
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	/* Operadores Relacionais */
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	/* Operadores Lógicos */
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	/* Operadores Unários */
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)
}
