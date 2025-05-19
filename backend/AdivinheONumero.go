package main

import (
	"fmt" // Entrada/saída formatada
	"math/rand" // Para geração de números aleatórios
	"strconv" // Para conversão de string para números
	"time" // Para obter o tempo atual, usado como semente aleatória
)

func JogoDeAdivinhacao() {
	// Inicializa o gerador de números aleatórios com uma semente baseada no tempo
	rand.Seed(time.Now().UnixNano())

	// Gera um número aleatório entre 1 e 30
	numeroSecreto := rand.Intn(10) + 1

	var tentativas int = 0 // Conta quantas tentativas o jogador fez
	var palpiteStr string // Armazena a entrada do usuário como texto
	var palpite int // Armazena o palpite convertido para número
	var err error // Captura possíveis erros durante a conversão

	fmt.Println("\n=== ADIVINHE O NÚMERO ===")
	fmt.Println("Estou pensando em um número entre 1 e 10")

	// Loop principal do jogo
	for {
		tentativas++ // Loop infinito até que seja explicitamente interrompido

		fmt.Print("Qual é o seu palpite? ")
		fmt.Scanln(&palpiteStr)

		// Converte o palpite para número
		palpite, err = strconv.Atoi(palpiteStr)
		if err != nil {
			fmt.Println("Por favor, digite apenas números!")
			tentativas--
			continue // Usado para pular o resto do loop e iniciar uma nova rodada
		}

		// Lógica principal do jogo
		if palpite < numeroSecreto {
			fmt.Println("Tente um número MAIOR!")
		} else if palpite > numeroSecreto {
			fmt.Println("Tente um número MENOR!")
		} else {
			// Acertou!
			fmt.Printf("\nPARABÉNS! Você acertou em %d tentativas!\n", tentativas)
			break
		}
	}
}

func AdivinheONumero() {
	JogoDeAdivinhacao()
}
