package main

import (
	"fmt"

	"github.com/badoux/checkmail"
	"github.com/jailtonjunior94/fundamentos/auxiliar"
)

func main() {
	auxiliar.Escrever()
	auxiliar.Escrever2()

	err := checkmail.ValidateFormat("jailton.junior94@outlook.com")
	fmt.Println(err)

	fmt.Println("Escrevendo no pacote main")
}
