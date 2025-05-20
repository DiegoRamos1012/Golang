package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// ProcessarOponente lida com a interação quando o jogador encontra um oponente
func ProcessarOponente(oponenteEncontrado string, inventario []string) (bool, []string) {
	continuaJogando := true

	// Mapeamento de oponentes para seus itens de counter
	oponenteCounters := map[string]string{
		"O Predador Alfa":             "Oitão",
		"O Pinscher":                  "Bicho de Pelúcia",
		"10 mil abelhas perigosas!":   "Máscara de Urso",
		"Árvore de maças envenenadas": "", // Não tem counter específico
	}

	// Verifica se o jogador tem o counter para esse oponente
	counterItem := oponenteCounters[oponenteEncontrado]
	temCounter := false
	counterIndex := -1

	if counterItem != "" {
		for i, item := range inventario {
			if item == counterItem {
				temCounter = true
				counterIndex = i
				break
			}
		}
	}

	// Verifica qual oponente é
	if oponenteEncontrado == "O Predador Alfa" {
		fmt.Println("\nVocê encontrou um patinho amarelo super fofo! O que deseja fazer?")
		fmt.Println("1 - Ignorar e seguir caminho")
		fmt.Println("2 - Acariciar o patinho")

		if temCounter {
			fmt.Printf("3 - Usar %s\n", counterItem)
		}

		var escolhaPato string
		var opcoes string
		if temCounter {
			opcoes = "1, 2, ou 3"
		} else {
			opcoes = "1 ou 2"
		}
		fmt.Printf("Escolha (%s): ", opcoes)
		fmt.Scanln(&escolhaPato)

		if escolhaPato == "1" {
			fmt.Println("Você ignora o pato e segue seu caminho.")
			fmt.Println("O patinho amarelo observa você passar e vai embora calmamente, muito bonitinho...")
			return continuaJogando, inventario
		} else if temCounter && escolhaPato == "3" {
			fmt.Println("\nVocê saca rapidamente o Oitão e dispara! O pato explode em uma nuvem de penas!")

			// Remove o item do inventário
			inventario = append(inventario[:counterIndex], inventario[counterIndex+1:]...)
			return continuaJogando, inventario
		} else {
			fmt.Println("\nAo acariciar o pato, você percebe que sua boca se abre revelando fileiras de dentes afiados!")
			fmt.Println("Era uma armadilha! O Predador Alfa disfarçado de pato te ataca!")

			// Verifica se o jogador tem o Oitão no inventário
			if temCounter {
				fmt.Println("Por sorte, você tem um Oitão no inventário e consegue se defender!")
				fmt.Println("Deseja usar o Oitão agora? (S/N)")

				var usarOitao string
				fmt.Scanln(&usarOitao)

				if usarOitao == "S" || usarOitao == "s" {
					fmt.Println("\nVocê saca rapidamente o Oitão e dispara! O pato explode em uma nuvem de penas!")

					// Remove o item do inventário
					inventario = append(inventario[:counterIndex], inventario[counterIndex+1:]...)
					return continuaJogando, inventario
				} else {
					fmt.Println("\nVocê decide não usar o Oitão. O Predador Alfa te devora! GAME OVER!")
					continuaJogando = false
				}
			} else {
				fmt.Println("\nO Predador Alfa te devora em um único bocado! GAME OVER!")
				continuaJogando = false
			}
		}
	} else if oponenteEncontrado == "Árvore de maças envenenadas" {
		fmt.Println("\nVocê encontrou uma árvore com maçãs deliciosas. O que deseja fazer?")
		fmt.Println("1 - Comer uma maçã")
		fmt.Println("2 - Ignorar e seguir caminho")

		var escolhaArvore string
		fmt.Print("Escolha (1 ou 2): ")
		fmt.Scanln(&escolhaArvore)

		if escolhaArvore == "2" {
			fmt.Println("Você decide ignorar a árvore e segue seu caminho.")
			return continuaJogando, inventario
		} else {
			fmt.Println("\nVocê comeu uma maçã. Infelizmente, eram envenenadas! GAME OVER!")
			continuaJogando = false
		}
	} else if oponenteEncontrado == "O Pinscher" {
		fmt.Println("\nUm pequeno cachorro aparece latindo furiosamente. O que deseja fazer?")
		fmt.Println("1 - Tentar passar correndo")
		fmt.Println("2 - Ficar parado e encarar")

		if temCounter {
			fmt.Printf("3 - Usar %s\n", counterItem)
		}

		var escolhaPinscher string
		var opcoes string
		if temCounter {
			opcoes = "1, 2, ou 3"
		} else {
			opcoes = "1 ou 2"
		}
		fmt.Printf("Escolha (%s): ", opcoes)
		fmt.Scanln(&escolhaPinscher)

		if temCounter && escolhaPinscher == "3" {
			fmt.Println("\nVocê joga o Bicho de Pelúcia! O Pinscher treme-treme fica entretido e você passa em segurança!")

			// Remove o item do inventário
			inventario = append(inventario[:counterIndex], inventario[counterIndex+1:]...)
			return continuaJogando, inventario
		} else {
			fmt.Println("Você subestima o cachorrinho. Tolo...")
			fmt.Println("Ele começa a latir descontroladamente. Seus latidos ensurdecedores fazem seus tímpanos explodirem enquanto ele devora suas pernas!")
			fmt.Println("GAME OVER! O Pinscher te devorou. Na sua lápide estará escrito: 'Morreu para um cachorro de 2kg'")
			continuaJogando = false
		}
	} else if oponenteEncontrado == "10 mil abelhas perigosas!" {
		fmt.Println("\nVocê ouve um zumbido que rapidamente se torna ensurdecedor.")
		fmt.Println("Uma nuvem negra de abelhas furiosas aparece no horizonte e avança na sua direção!")

		if temCounter {
			fmt.Println("1 - Tentar fugir")
			fmt.Printf("2 - Usar %s\n", counterItem)

			var escolhaAbelhas string
			fmt.Print("Escolha (1 ou 2): ")
			fmt.Scanln(&escolhaAbelhas)

			if escolhaAbelhas == "2" {
				fmt.Println("Você coloca a Máscara de Urso! As abelhas se assustam e fogem!")

				// Remove o item do inventário
				inventario = append(inventario[:counterIndex], inventario[counterIndex+1:]...)
				return continuaJogando, inventario
			}
			// Se escolher 1 ou qualquer outra coisa, continua para o game over
		}

		fmt.Println("Em segundos, milhares de ferrões penetram sua pele de todos os lados.")
		fmt.Println("A última coisa que você vê antes de desmaiar é seu corpo transformado em um cacto humano.")
		fmt.Println("GAME OVER! Você virou um cacto de tanto ferrão espetado. Se alguém encontrar seu corpo,")
		fmt.Println("vai achar que é uma nova espécie de planta do deserto.")
		continuaJogando = false
	} else {
		fmt.Printf("\nOh não! Você encontrou %s! GAME OVER!\n", oponenteEncontrado)
		continuaJogando = false
	}

	return continuaJogando, inventario
}

// GerarEventoAleatorio gera um evento aleatório no jogo
func GerarEventoAleatorio() {
	eventosAleatorios := []string{
		"Um pássaro pousa na sua cabeça e te confunde com uma árvore por 1 turno",
		"Você pisa em uma banana e escorrega para uma posição aleatória do mapa",
		"Você encontra um espelho mágico e fica admirando seu cabelo perfeito por 1 turno",
		"Uma sereia aparece e te convida para um chá. Depois de recusar, ela fica emburrada e vai embora",
		"Você achou sinal de internet, suficiente pra carregar os reels do instagram. Você assiste um pouco e percebe que ficou uma hora aqui, que loucura, né?",
	}
	fmt.Println(eventosAleatorios[rand.Intn(len(eventosAleatorios))])
}

// VerificarAoRedor retorna se há um tesouro ou oponente nas células adjacentes
func VerificarAoRedor(jogadorX, jogadorY int, mapa [][]string) (bool, int, int, string) {
	// Verifica células adjacentes
	direcoes := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Norte, Sul, Oeste, Leste

	for _, dir := range direcoes {
		novoX, novoY := jogadorX+dir[0], jogadorY+dir[1]

		// Verifica se está dentro dos limites
		if novoX >= 0 && novoX < 5 && novoY >= 0 && novoY < 5 {
			if mapa[novoX][novoY] == "O" {
				return true, novoX, novoY, "O"
			} else if mapa[novoX][novoY] == "T" {
				return true, novoX, novoY, "T"
			}
		}
	}
	return false, 0, 0, ""
}

// ExibirInventario exibe o inventário do jogador
func ExibirInventario(inventario []string) {
	if len(inventario) > 0 {
		fmt.Println("Seu inventário:", strings.Join(inventario, ", "))
	}
}

// AplicarEfeitoItem verifica e aplica o efeito de um item em um alvo específico
func AplicarEfeitoItem(itemSelecionado, alvoTipo string, alvoX, alvoY, jogadorX, jogadorY int, mapa [][]string, mapaOponentes map[string]string) (novoX, novoY int, posicaoOriginal string, efeitoAplicado bool, continuaJogando bool) {
	continuaJogando = true
	efeitoAplicado = false
	novoX, novoY = jogadorX, jogadorY
	posicaoOriginal = mapa[jogadorX][jogadorY]

	// Determina qual oponente está na posição alvo
	chave := fmt.Sprintf("%d,%d", alvoX, alvoY)
	oponenteAtual := mapaOponentes[chave]

	switch {
	// Botas Saltitantes contra qualquer oponente - com chance de falha
	case itemSelecionado == "Botas Saltitantes" && alvoTipo == "O":
		// 70% de chance de sucesso, 30% de falha
		if rand.Intn(10) < 7 {
			fmt.Println("Você usa as Botas Saltitantes e pula por cima do oponente. Que sorte!")
			// Encontra posição vazia atrás do oponente na mesma direção
			novoX = alvoX + (alvoX - jogadorX)
			novoY = alvoY + (alvoY - jogadorY)

			if novoX >= 0 && novoX < 5 && novoY >= 0 && novoY < 5 && mapa[novoX][novoY] == "~~" {
				// Restaura a posição original
				mapa[jogadorX][jogadorY] = posicaoOriginal

				// Move o jogador para a nova posição
				posicaoOriginal = mapa[novoX][novoY]

				// Primeiro exibe o mapa
				exibirMapa(mapa)

				// Depois exibe a mensagem
				fmt.Println("Você saltou para uma nova posição!")
				efeitoAplicado = true
			} else {
				fmt.Println("Não é possível saltar para aquela direção!")
			}
		} else {
			fmt.Println("Ops! Você pulou muito alto e acabou alcançando o céu! Infelizmente, as botas não vão amenizar sua queda...GAME OVER!")
			efeitoAplicado = true
			continuaJogando = false
		}

	// Oitão específico contra o Predador Alfa
	case itemSelecionado == "Oitão" && alvoTipo == "O":
		if oponenteAtual == "O Predador Alfa" {
			fmt.Println("Você saca o Oitão e dispara contra o patinho amarelo! POW!")
			fmt.Println("O pato explode em uma nuvem de penas amarelas. Pato safado mereceu!")
			mapa[alvoX][alvoY] = "~~"
			efeitoAplicado = true
		} else if oponenteAtual == "O Pinscher" {
			fmt.Println("Você tenta usar o Oitão contra o Pinscher, mas ele é muito ágil e desvia dos tiros!")
			fmt.Println("O cachorrinho furioso avança e abocanha sua perna. GAME OVER!")
			efeitoAplicado = true
			continuaJogando = false
		} else if oponenteAtual == "10 mil abelhas perigosas!" {
			fmt.Println("Você atira contra as abelhas, mas são muitas! Seu tiro elimina apenas umas 10 delas...")
			fmt.Println("As outras 9.990 abelhas ficam furiosas e atacam juntas. GAME OVER!")
			efeitoAplicado = true
			continuaJogando = false
		} else if oponenteAtual == "Árvore de maças envenenadas" {
			fmt.Println("Você atira contra a árvore... ela não se move, claro, é uma árvore.")
			fmt.Println("Mas uma maçã cai com o impacto, você a pega e come sem pensar. Era envenenada! GAME OVER!")
			efeitoAplicado = true
			continuaJogando = false
		} else {
			fmt.Println("O Oitão não teve efeito contra este oponente!")
		}

	case itemSelecionado == "Máscara de Urso" && alvoTipo == "O":
		if oponenteAtual == "10 mil abelhas perigosas!" {
			fmt.Println("Você coloca a Máscara de Urso! As abelhas se assustam e fogem!")
			mapa[alvoX][alvoY] = "~~"
			efeitoAplicado = true
		} else if oponenteAtual == "O Predador Alfa" {
			fmt.Println("Você coloca a Máscara de Urso e se aproxima do patinho amarelo...")
			fmt.Println("Para sua surpresa, o pato se transforma no Predador Alfa e começa a lamber os beiços.")
			fmt.Println("Infelizmente, você descobriu que ele gosta mais de ursos do que de humanos. Você foi de topo da cadeia alimentar pra café da manhã. GAME OVER!")
			efeitoAplicado = true
			continuaJogando = false
		} else if oponenteAtual == "O Pinscher" {
			fmt.Println("Você coloca a Máscara de Urso. O Pinscher olha pra você, ri e ataca mesmo assim.")
			fmt.Println("Aparentemente ele não tem medo de ursos. GAME OVER!")
			efeitoAplicado = true
			continuaJogando = false
		} else if oponenteAtual == "Árvore de maças envenenadas" {
			fmt.Println("Você coloca a Máscara de Urso perto da árvore, mas nada acontece...")
			fmt.Println("Você acaba pegando uma maçã para comer enquanto espera alguma reação. Era envenenada! GAME OVER!")
			efeitoAplicado = true
			continuaJogando = false
		} else {
			fmt.Println("A Máscara de Urso não teve efeito contra este oponente!")
		}

	default:
		fmt.Println("Este item não tem efeito aqui!")
	}

	return novoX, novoY, posicaoOriginal, efeitoAplicado, continuaJogando
}

// ProcessarTesouro lida com a interação quando o jogador encontra um tesouro
func ProcessarTesouro(tesouroEscolhido string, turnos int) bool {
	// 70% de chance do Guardião do Tesouro aparecer
	if rand.Intn(10) < 7 {
		fmt.Println("\nQuando você está prestes a pegar o tesouro, um ser misterioso aparece!")
		fmt.Println("'Sou o Guardião do Tesouro! Para levar esta relíquia, terá que me vencer em um duelo mortal...no Pedra, Papel, Tesoura!'")
		fmt.Println("O guardião sorri maliciosamente. 'Se você perder... bem, digamos que o tesouro está com fome.'")

		// Jogo de Pedra, Papel, Tesoura
		fmt.Println("\n=== PEDRA, PAPEL, TESOURA ===")
		fmt.Println("1 - Pedra")
		fmt.Println("2 - Papel")
		fmt.Println("3 - Tesoura")

		var escolhaJogador int
		fmt.Print("Faça sua escolha (1-3): ")
		fmt.Scanln(&escolhaJogador)

		// Valida a entrada
		if escolhaJogador < 1 || escolhaJogador > 3 {
			fmt.Println("Escolha inválida! O Guardião ri da sua confusão...")
			escolhaJogador = rand.Intn(3) + 1 // Escolha aleatória se input for inválido
			fmt.Printf("Ele decide que você escolheu %d!\n", escolhaJogador)
		}

		// Escolha do guardião (1-Pedra, 2-Papel, 3-Tesoura)
		escolhaGuardiao := rand.Intn(3) + 1

		// Converte números para nomes
		opcoes := []string{"", "Pedra", "Papel", "Tesoura"}

		fmt.Printf("\nVocê escolheu: %s\n", opcoes[escolhaJogador])
		fmt.Printf("O Guardião escolheu: %s\n", opcoes[escolhaGuardiao])

		// Determina o vencedor
		// Empate: escolhas iguais
		// Vitória do jogador: (Pedra vs Tesoura) OU (Papel vs Pedra) OU (Tesoura vs Papel)
		if escolhaJogador == escolhaGuardiao {
			fmt.Println("\nEMPATE! O Guardião parece frustrado...")
			fmt.Println("'Isso raramente acontece... vamos de novo!'")

			// Nova rodada em caso de empate
			escolhaJogador = rand.Intn(3) + 1
			escolhaGuardiao = rand.Intn(3) + 1

			fmt.Printf("\nVocê escolheu: %s\n", opcoes[escolhaJogador])
			fmt.Printf("O Guardião escolheu: %s\n", opcoes[escolhaGuardiao])
		}

		// Verifica novamente o resultado
		if escolhaJogador == escolhaGuardiao {
			fmt.Println("\nOutro EMPATE! O Guardião suspira profundamente.")
			fmt.Println("'Ah, de novo? Cansei, pega logo esse negócio, vou me aposentar'")
			fmt.Printf("\nPARABÉNS! Você encontrou o... %s após %d turnos!\n", tesouroEscolhido, turnos)
			return false
		} else if (escolhaJogador == 1 && escolhaGuardiao == 3) ||
			(escolhaJogador == 2 && escolhaGuardiao == 1) ||
			(escolhaJogador == 3 && escolhaGuardiao == 2) {
			fmt.Println("\nVOCÊ VENCEU! O Guardião aceita a derrota com dignidade.")
			fmt.Println("'Um acordo é um acordo. O tesouro é seu, aventureiro.'")
			fmt.Printf("\nPARABÉNS! Você encontrou o... %s após %d turnos!\n", tesouroEscolhido, turnos)
			return false
		} else {
			fmt.Println("\nVOCÊ PERDEU! O Guardião solta uma gargalhada maléfica.")
			fmt.Println("'Eu avisei sobre a fome do tesouro... agora observe!'")
			fmt.Println("Diante dos seus olhos horrorizados, o tesouro desenvolve fileiras de dentes afiados.")

			if tesouroEscolhido == "Poder da Amizade" {
				fmt.Println("O Poder da Amizade sorri para você... antes de devorar sua face! Baú difícil de fazer amizade")
			} else if tesouroEscolhido == "Rubi de Sangue" {
				fmt.Println("O Rubi de Sangue fica cada vez mais vermelho enquanto suga todo o seu sangue!")
			} else if tesouroEscolhido == "Vaso de Jade" {
				fmt.Println("O Vaso de Jade se transforma numa armadilha mortal, prendendo você dentro e fechando a tampa!")
			} else {
				fmt.Println("O Pergaminho Ancião se desenrola e envolve seu corpo como uma múmia, sugando sua alma! Que dramático!")
			}

			fmt.Println("GAME OVER! Quem diria que tesouros poderiam ser tão perigosos? 'A ganância que te move é a mesma que te mata'")
			return false
		}
	} else {
		// 30% de chance de não encontrar o Guardião
		fmt.Printf("\nPARABÉNS! Depois de muito trabalho árduo, você encontrou o... %s após %d turnos!\n", tesouroEscolhido, turnos)
		return false
	}
}
