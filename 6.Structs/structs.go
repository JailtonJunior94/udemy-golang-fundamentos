package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var usuario1 usuario
	usuario1.nome = "Jailton"
	usuario1.idade = 26

	fmt.Println(usuario1)

	usuario2 := usuario{"Stefany", 26, endereco{"José Pontes", 156}}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Antony"}
	fmt.Println(usuario3)
}
