package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2
	return media >= 6
}

func main() {
	defer funcao1()
	// DEFER = ADIAR
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
