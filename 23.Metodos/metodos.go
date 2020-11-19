package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.maiorDeIdade() >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Jailton", 26}
	usuario1.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
