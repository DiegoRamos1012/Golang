package main

import (
	"errors"
	"fmt"
	"strconv"
)

func VerificarPar(num int) (int, error) {
	if num%2 == 0 {
		fmt.Println("Este número é par")
		return num, nil
	} else {
		fmt.Println("Este número é ímpar")
		return num, errors.New("Este número é ímpar")
	}
}

func ParOuImpar() {
	var inputNum string
	var num int64
	var err error

	fmt.Println("\n==== Verificador de Par ou Ímpar ====")
	fmt.Printf("Insira o número a ser verificado: ")
	fmt.Scanln(&inputNum)

	num, err = strconv.ParseInt(inputNum, 10, 64)
	if err != nil {
		fmt.Println("Erro: Número inválido")
		return
	}
	_, err = VerificarPar(int(num))
}
