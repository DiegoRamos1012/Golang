package main

import (
	"fmt"
	"os"
)

func ChecarNumero() {
	var num int64

	fmt.Print("Insira um número: ")

	// Captura o retorno de Scanln para verificar erros
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Erro na leitura. Por favor, digite um número válido.")
		return
	}

	// Inclusão dos limites 1 e 10 na verificação
	if num >= 1 && num <= 10 {
		fmt.Printf("%v está entre 1 e 10\n", num)
	} else {
		fmt.Printf("%v não está entre 1 e 10\n", num)
	}
}

// Versão alternativa que retorna um booleano e é mais reutilizável
func VerificarEntreUmEDez(num int64) bool {
	return num >= 1 && num <= 10
}

// Função principal para demonstração
func main() {
	ChecarNumero()

	// Para evitar que o console feche imediatamente em ambientes Windows
	fmt.Println("\nPressione Enter para sair...")
	fmt.Scanln()
	os.Exit(0)
}
