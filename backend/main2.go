package main

import (
	// O pacote "errors" permite criar e manipular erros personalizados
	"errors"
	// O pacote "fmt" fornece funções para formatação e entrada/saída de dados
	"fmt"
)

func dividir(a, b float64) (float64, error) { //float64 são números decimais de precisão dupla
	if b == 0 {
		return 0, errors.New("Não é possível dividir por zero")		
	}
	return a / b, nil // nil significa "nenhum erro ocorreu"
}

func exemploErros() {
	var a, b float64

	fmt.Print("Digite o primeiro número: ")
	// Lê o valor digitado pelo usuário e armazena na variável "a"
    // O "&" obtém o endereço de memória da variável para que Scan possa modificá-la
	fmt.Scan(&a)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&b)

	// Chama a função dividir com os valores informados e captura os retornos
    // - resultado: contém o valor da divisão (se bem-sucedida)
    // - err: contém informações sobre o erro (se ocorrer)
	resultado, err := dividir(a, b)
	if err != nil {
		// Se houver erro, exibe a mensagem de erro para o usuário
		fmt.Println("Erro:", err)
	} else {
		
	// %g é um especificador de formato que remove zeros desnecessários à direita
    fmt.Printf("Resultado de %g / %g = %g\n", a, b, resultado)
	}	
}

