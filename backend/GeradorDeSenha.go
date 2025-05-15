package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GerarSenha cria uma senha aleatória com o comprimento especificado
func GerarSenha(comprimento int) (string, error) {
	// Define os caracteres possíveis
	const caracteres = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?"

	// Prepara um slice de bytes para armazenar a senha
	senha := make([]byte, comprimento)

	// Preenche cada posição da senha
	for i := range comprimento {
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
	senha, err := GerarSenha(12) // Gera uma senha de 12 caracteres
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Senha gerada:", senha)
}
