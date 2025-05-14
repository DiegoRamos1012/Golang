package main

import (
	// O pacote "errors" permite criar e manipular erros personalizados
	"errors"
	// O pacote "fmt" fornece funções para formatação e entrada/saída de dados
	"fmt"
	"strconv"
)

func somar(a,b float64) (float64, error) {
	return a + b, nil	
}

func subtrair (a, b float64) (float64, error) {
	return a - b, nil
}

func multiplicar (a, b float64) (float64, error) {
	return a * b, nil
}

func dividir(a, b float64) (float64, error) { //float64 são números decimais de precisão dupla
	if b == 0 {
		return 0, errors.New("Não é possível dividir por zero")		
	}
	return a / b, nil // nil significa "nenhum erro ocorreu"
}

func calculadora() {
    var inputA, inputB, escolha string
    var a, b float64
    var err error

    fmt.Println("\n==== Calculadora Simples ====")
    fmt.Println("1 - Soma")
    fmt.Println("2 - Subtração")
    fmt.Println("3 - Multiplicação")
    fmt.Println("4 - Divisão")
    fmt.Print("\nEscolha uma operação (1-4): ")
    fmt.Scanln(&escolha)

    fmt.Print("Digite o primeiro número: ")
    fmt.Scanln(&inputA)
    
    // Converte o primeiro número
    a, err = strconv.ParseFloat(inputA, 64)
    if err != nil {
        fmt.Println("Erro: O primeiro valor não é um número válido")
        return
    }

    fmt.Print("Digite o segundo número: ")
    fmt.Scanln(&inputB)
    
    // Converte o segundo número
    b, err = strconv.ParseFloat(inputB, 64)
    if err != nil {
        fmt.Println("Erro: O segundo valor não é um número válido")
        return
    }

    // Executa a operação escolhida
    switch escolha {
    case "1":
        resultado, _ := somar(a, b)
        fmt.Printf("Soma: %g + %g = %g\n", a, b, resultado)
    case "2":
        resultado, _ := subtrair(a, b)
        fmt.Printf("Subtração: %g - %g = %g\n", a, b, resultado)
    case "3":
        resultado, _ := multiplicar(a, b)
        fmt.Printf("Multiplicação: %g × %g = %g\n", a, b, resultado)
    case "4":
        resultado, err := dividir(a, b)
        if err != nil {
            fmt.Println("Erro:", err)
        } else {
            fmt.Printf("Divisão: %g ÷ %g = %g\n", a, b, resultado)
        }
    default:
        fmt.Println("Opção inválida!")
    }
}

