package main

import "fmt"

func alterarValor(numero *int) {
	*numero = 42
}

func Ponteiros() {
	x := 10
	fmt.Println("Antes:", x)

	alterarValor(&x)
	fmt.Println("Depois:", x)
}

