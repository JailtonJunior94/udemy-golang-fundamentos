package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough // Faz cai na próxima validação
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia)
}
