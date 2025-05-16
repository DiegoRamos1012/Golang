package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

// GerarSenha cria uma senha aleatória com o comprimento especificado
func GerarSenha(comprimento int) (string, error) {
	// Define os caracteres possíveis
	const caracteres = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?"

	// Prepara um slice de bytes para armazenar a senha
	senha := make([]byte, comprimento)

	// Preenche cada posição da senha
	for i := range senha { // Note: aqui tinha um erro, deve iterar sobre senha, não comprimento
		// Gera um número aleatório entre 0 e len(caracteres)-1
		indice, err := rand.Int(rand.Reader, big.NewInt(int64(len(caracteres))))
		if err != nil {
			return "", err
		}
		// Seleciona o caractere correspondente ao índice gerado
		senha[i] = caracteres[indice.Int64()]
	}

	return string(senha), nil
}

func GeradorDeSenha() {
	inicio := time.Now()
	var wg sync.WaitGroup

	// Número de senhas a gerar
	numSenhas := 5

	// Slice para armazenar resultados
	senhas := make([]string, numSenhas)

	// Iniciar múltiplas goroutines para gerar senhas concorrentemente
	for i := 0; i < numSenhas; i++ {
		wg.Add(1)

		// O "go" inicia uma nova goroutine
		go func(index int) {
			defer wg.Done()

			// Gerar senha com tamanho baseado no índice
			tamanho := 10 + index
			senha, err := GerarSenha(tamanho)
			if err != nil {
				fmt.Printf("Erro na goroutine %d: %v\n", index, err)
				return
			}

			// Armazenar resultado
			senhas[index] = senha
			fmt.Printf("Goroutine %d gerou senha: %s\n", index, senha)
		}(i)
	}

	// Esperar todas as goroutines terminarem
	wg.Wait()

	fmt.Println("\nTodas as senhas geradas:")
	for i, senha := range senhas {
		fmt.Printf("Senha %d (%d caracteres): %s\n", i, 10+i, senha)
	}

	fmt.Printf("\nTempo total: %v\n", time.Since(inicio))
}
