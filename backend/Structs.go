package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func Structs() {
	pessoa := Pessoa{"Maria", 30}
	fmt.Println("Nome: ", pessoa.Nome)
	fmt.Println("Nome: ", pessoa.Idade)
}