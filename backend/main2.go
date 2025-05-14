package main

import (
	"errors"
	"fmt"
)

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Não é possível dividir por zero")		
	}
	return a / b, nil
}

func exemploErros() {
	var a, b float64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&a)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&b)

	resultado, err := dividir(a, b)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Resultado de %.2f / %.2f = %.2f\n", a, b, resultado)
	}
}
