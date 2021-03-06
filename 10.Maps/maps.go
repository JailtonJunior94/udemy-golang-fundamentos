package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Jailton",
		"sobrenome": "Junior",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")

	fmt.Println(usuario2)
}
