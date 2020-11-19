package main

import "fmt"

func generica(param interface{}) {
	fmt.Println(param)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
}
