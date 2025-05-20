package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

func IlhaDoTesouro() {
	// Inicializa a semente aleatória
	rand.Seed(time.Now().UnixNano())

	// Define os tesouros e oponentes (corrigindo o formato do array)
	tesouros := []string{"Rubi de Sangue", "Vaso de Jade", "Pergaminho Ancião", "Poder da Amizade"}
	oponentes := []string{"O Predador Alfa", "O Pinscher", "Árvore de maças envenenadas", "10 mil abelhas perigosas!"}
	itensMagicos := []string{"Botas Saltitantes", "Oitão", "Máscara de Urso", "Bicho de Pelúcia"}

	// Cria um mapa vazio 5x5
	mapa := make([][]string, 5)
	for i := range mapa {
		mapa[i] = make([]   string, 5)
		// Inicializa todas as posições como vazias
		for j := range mapa[i] {
			mapa[i][j] = "~~" // Representando água ou área vazia
		}
	}

	// Posiciona o tesouro aleatoriamente
	tesouroEscolhido := tesouros[rand.Intn(len(tesouros))]
	tesouroX, tesouroY := rand.Intn(5), rand.Intn(5)
	mapa[tesouroX][tesouroY] = "T" // T representa tesouro

	// Posiciona alguns oponentes aleatoriamente (2 oponentes)
	oponentesEncontrados := make([]string, 2)
	for i := 0; i < 2; i++ {
		oponentesEncontrados[i] = oponentes[rand.Intn(len(oponentes))]

		// Encontra uma posição vazia
		var oponenteX, oponenteY int
		for {
			oponenteX, oponenteY = rand.Intn(5), rand.Intn(5)
			if mapa[oponenteX][oponenteY] == "~~" {
				break // Encontrou uma posição vazia
			}
		}

		mapa[oponenteX][oponenteY] = "O" // O representa oponente
	}

	// Posiciona itens mágicos aleatoriamente (3 itens)
	itensColocados := make([]string, 3)
	for i := range 3 {
		itensColocados[i] = itensMagicos[rand.Intn(len(itensMagicos))]

		// Encontra uma posição vazia
		var itemX, itemY int
		for {
			itemX, itemY = rand.Intn(5), rand.Intn(5)
			if mapa[itemX][itemY] == "~~" {
				break // Encontrou uma posição vazia
			}
		}

		mapa[itemX][itemY] = "I" // I representa item mágico
	}

	// Define a posição inicial do jogador (sempre no canto superior esquerdo)
	jogadorX, jogadorY := 0, 0

	// Guarda a posição original do que está na posição do jogador
	posicaoOriginal := mapa[jogadorX][jogadorY]
	mapa[jogadorX][jogadorY] = "P" // P representa o jogador

	// Inventário do jogador (itens coletados)
	inventario := []string{}

	// Apresentação do jogo
	fmt.Println("\n=== ILHA DO TESOURO ===")
	fmt.Println("Bem vindo a Ilha Oculta. Existe um tesouro escondido aqui, encontre-o!")
	fmt.Println("Legenda: P = Você, ~~ = Área vazia")
	fmt.Println("Existem itens mágicos, tesouros e oponentes escondidos!")
	fmt.Println("Use: 'N' (Norte), 'S' (Sul), 'L' (Leste), 'O' (Oeste) para se mover")
	fmt.Println("Use: 'U' (Usar) para utilizar um item do seu inventário")

	// Exibe o mapa inicial
	exibirMapa(mapa)

	// Variáveis de controle do jogo
	jogando := true
	turnos := 0
	var escolha string

	// Loop principal do jogo
	for jogando {
		turnos++

		// Exibe o inventário se tiver itens
		if len(inventario) > 0 {
			fmt.Println("Seu inventário:", strings.Join(inventario, ", "))
		}

		fmt.Print("O que deseja fazer? (N/S/L/O para mover, I para usar item): ")
		fmt.Scanln(&escolha)

		// Verifica se o jogador quer usar um item
		if escolha == "I" || escolha == "i" {
			if len(inventario) == 0 {
				fmt.Println("Você não tem itens no inventário!")
				continue
			}

			fmt.Println("Qual item deseja usar?")
			for i, item := range inventario {
				fmt.Printf("%d - %s\n", i+1, item)
			}

			var itemEscolhido int
			fmt.Print("Escolha o número do item: ")
			fmt.Scanln(&itemEscolhido)

			if itemEscolhido < 1 || itemEscolhido > len(inventario) {
				fmt.Println("Escolha inválida!")
				continue
			}

			itemSelecionado := inventario[itemEscolhido-1]

			// Verifica se está perto de um oponente ou tesouro
			temAlgo := false

			// Verifica células adjacentes
			direcoes := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Norte, Sul, Oeste, Leste
			var alvoX, alvoY int
			var alvoTipo string

			for _, dir := range direcoes {
				novoX, novoY := jogadorX+dir[0], jogadorY+dir[1]

				// Verifica se está dentro dos limites
				if novoX >= 0 && novoX < 5 && novoY >= 0 && novoY < 5 {
					if mapa[novoX][novoY] == "O" {
						temAlgo = true
						alvoX, alvoY = novoX, novoY
						alvoTipo = "O"
						break
					} else if mapa[novoX][novoY] == "T" {
						temAlgo = true
						alvoX, alvoY = novoX, novoY
						alvoTipo = "T"
						break
					}
				}
			}

			if !temAlgo {
				fmt.Println("Não há nada por perto para usar este item!")
				continue
			}

			// Define o efeito do item de acordo com o alvo
			var efeitoAplicado bool = false

			switch {
			// Botas Saltitantes contra qualquer oponente - com chance de falha
			case itemSelecionado == "Botas Saltitantes" && alvoTipo == "O":
				// 70% de chance de sucesso, 30% de falha
				if rand.Intn(10) < 7 {
					fmt.Println("Você usa as Botas Saltitantes e pula por cima do oponente. Que sorte!")
					// Encontra posição vazia atrás do oponente na mesma direção
					novoX := alvoX + (alvoX - jogadorX)
					novoY := alvoY + (alvoY - jogadorY)

					if novoX >= 0 && novoX < 5 && novoY >= 0 && novoY < 5 && mapa[novoX][novoY] == "~~" {
						// Restaura a posição original
						mapa[jogadorX][jogadorY] = posicaoOriginal

						// Move o jogador para a nova posição
						jogadorX, jogadorY = novoX, novoY
						posicaoOriginal = mapa[jogadorX][jogadorY]
						mapa[jogadorX][jogadorY] = "P"

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
					jogando = false
				}

				// Oitão específico contra o Predador Alfa
			case itemSelecionado == "Oitão" && alvoTipo == "O":
				// Determina qual oponente encontrou
				oponenteAtual := ""
				for _, oponente := range oponentesEncontrados {
					if mapa[alvoX][alvoY] == "O" {
						oponenteAtual = oponente
						break
					}
				}

				if oponenteAtual == "O Predador Alfa" {
					fmt.Println("Você saca o Oitão e dispara contra o patinho amarelo! POW!")
					fmt.Println("O pato explode em uma nuvem de penas amarelas. Pato safado mereceu")
					mapa[alvoX][alvoY] = "~~"
					efeitoAplicado = true
				} else {
					fmt.Println("O Oitão não teve efeito contra este oponente!")
				}

				// Máscara de Urso específica contra as abelhas
			case itemSelecionado == "Máscara de Urso" && alvoTipo == "O":
				// Determina qual oponente encontrou
				oponenteAtual := ""
				for _, oponente := range oponentesEncontrados {
					if mapa[alvoX][alvoY] == "O" {
						oponenteAtual = oponente
						break
					}
				}

				if oponenteAtual == "10 mil abelhas perigosas!" {
					fmt.Println("Você coloca a Máscara de Urso! As abelhas se assustam e fogem!")
					mapa[alvoX][alvoY] = "~~"
					efeitoAplicado = true
				} else if oponenteAtual == "O Predador Alfa" {
					fmt.Println("Você coloca a Máscara de Urso e se aproxima do patinho amarelo...")
					fmt.Println("Para sua surpresa, o pato se transforma no Predador Alfa e começa a lamber os beiços.")
					fmt.Println("Infelizmente, você descobriu que ele gosta mais de ursos do que de humanos. Você foi de topo da cadeia alimentar pra café da manhã. GAME OVER!")
					efeitoAplicado = true
					jogando = false
				} else {
					fmt.Println("A Máscara de Urso não teve efeito contra este oponente!")
				}

				// Bicho de Pelúcia específico contra o Pinscher
			case itemSelecionado == "Bicho de Pelúcia" && alvoTipo == "O":
				// Determina qual oponente encontrou
				oponenteAtual := ""
				for _, oponente := range oponentesEncontrados {
					if mapa[alvoX][alvoY] == "O" {
						oponenteAtual = oponente
						break
					}
				}

				if oponenteAtual == "O Pinscher" {
					fmt.Println("Você joga o Bicho de Pelúcia! O Pinscher treme-treme fica entretido e você passa em segurança!")
					mapa[alvoX][alvoY] = "~~"
					efeitoAplicado = true
				} else if oponenteAtual == "Árvore de maças envenenadas" {
					fmt.Println("Você tenta usar o Bicho de Pelúcia na árvore, mas acaba comendo uma maçã envenenada! GAME OVER!")
					efeitoAplicado = true
					jogando = false
				} else if oponenteAtual == "O Predador Alfa" {
					fmt.Println("Você joga o Bicho de Pelúcia para o patinho amarelo...")
					fmt.Println("Ele olha com curiosidade, então num movimento rápido engole o brinquedo inteiro!")
					fmt.Println("Parece que ele só viu o pelúcia como uma sobremesa, e você como o jantar")
					fmt.Println("Seus ossos estalam enquanto o pato te engole inteiro. GAME OVER!")
					efeitoAplicado = true
					jogando = false
				} else {
					fmt.Println("O Bicho de Pelúcia não teve efeito contra este oponente!")
				}

				// Bicho de Pelúcia contra o tesouro (mantido da versão anterior)
			case alvoTipo == "T":
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
						fmt.Printf("\nPARABÉNS! Você encontrou o...%s após %d turnos!\n", tesouroEscolhido, turnos)
						jogando = false
					} else if (escolhaJogador == 1 && escolhaGuardiao == 3) ||
						(escolhaJogador == 2 && escolhaGuardiao == 1) ||
						(escolhaJogador == 3 && escolhaGuardiao == 2) {
						fmt.Println("\nVOCÊ VENCEU! O Guardião aceita a derrota com dignidade.")
						fmt.Println("'Um acordo é um acordo. O tesouro é seu, aventureiro.'")
						fmt.Printf("\nPARABÉNS! Você encontrou o...%s após %d turnos!\n", tesouroEscolhido, turnos)
						jogando = false
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
						jogando = false
					}
				} else {
					// 30% de chance de não encontrar o Guardião
					fmt.Printf("\nPARABÉNS! Depois de muito trabalho árduo, você encontrou o...%s após %d turnos!\n", tesouroEscolhido, turnos)
					jogando = false
				}

			default:
				fmt.Println("Este item não tem efeito aqui!")
			}

			// Remove o item do inventário se foi usado
			if efeitoAplicado {
				inventario = slices.Delete(inventario, itemEscolhido-1, itemEscolhido)
			}

			// Exibe mapa atualizado se o jogo continua
			if jogando {
				exibirMapa(mapa)
			}

			continue
		}

		// Guarda posição anterior
		antigaX, antigaY := jogadorX, jogadorY

		// Processa o movimento
		switch escolha {
		case "N", "n":
			if jogadorX > 0 {
				jogadorX--
			} else {
				fmt.Println("Você não pode ir para o Norte. Está no limite do mapa!")
				continue
			}
		case "S", "s":
			if jogadorX < 4 {
				jogadorX++
			} else {
				fmt.Println("Você não pode ir para o Sul. Está no limite do mapa!")
				continue
			}
		case "L", "l":
			if jogadorY < 4 {
				jogadorY++
			} else {
				fmt.Println("Você não pode ir para o Leste. Está no limite do mapa!")
				continue
			}
		case "O", "o":
			if jogadorY > 0 {
				jogadorY--
			} else {
				fmt.Println("Você não pode ir para o Oeste. Está no limite do mapa!")
				continue
			}
		default:
			fmt.Println("Comando inválido! Use N, S, L, O para mover ou U para usar item.")
			continue
		}

		// Restaura o que estava na posição anterior
		mapa[antigaX][antigaY] = posicaoOriginal

		// Verifica o que há na nova posição
		posicaoOriginal = mapa[jogadorX][jogadorY]

		// Verifica se encontrou algo
		switch posicaoOriginal {
		case "T":
			fmt.Printf("\nPARABÉNS! Depois de muito trabalho árduo e sono ruim, Você encontrou o...%s após %d turnos!\n", tesouroEscolhido, turnos)
			jogando = false

		case "O":
			// Determina qual oponente encontrou
			var oponenteEncontrado string
			oponenteIndex := rand.Intn(len(oponentesEncontrados))
			oponenteEncontrado = oponentesEncontrados[oponenteIndex]

			// Verifica qual oponente é
			if oponenteEncontrado == "O Predador Alfa" {
				fmt.Println("\nVocê encontrou um patinho amarelo super fofo! O que deseja fazer?")
				fmt.Println("1 - Ignorar e seguir caminho")
				fmt.Println("2 - Acariciar o patinho")

				var escolhaPato string
				fmt.Print("Escolha (1 ou 2): ")
				fmt.Scanln(&escolhaPato)

				if escolhaPato == "1" {
					fmt.Println("Você ignora o pato e segue seu caminho.")
					fmt.Println("O patinho amarelo observa você passar e vai embora calmamente, muito bonitinho...")

					// Importante: muda o oponente para espaço vazio
					posicaoOriginal = "~~"
					mapa[jogadorX][jogadorY] = "P"

					// Exibe o mapa atualizado
					exibirMapa(mapa)
					continue
				} else {
					fmt.Println("\nAo acariciar o pato, você percebe que sua boca se abre revelando fileiras de dentes afiados!")
					fmt.Println("Era uma armadilha! O Predador Alfa disfarçado de pato te ataca!")

					// Verifica se o jogador tem o Oitão no inventário
					temOitao := false
					for _, item := range inventario {
						if item == "Oitão" {
							temOitao = true
							break
						}
					}

					if temOitao {
						fmt.Println("Por sorte, você tem um Oitão no inventário e consegue se defender!")
						fmt.Println("Deseja usar o Oitão agora? (S/N)")

						var usarOitao string
						fmt.Scanln(&usarOitao)

						if usarOitao == "S" || usarOitao == "s" {
							fmt.Println("\nVocê saca rapidamente o Oitão e dispara! O pato explode em uma nuvem de penas!")
							posicaoOriginal = "~~"
							mapa[jogadorX][jogadorY] = "P"

							// Remove o Oitão do inventário
							for i, item := range inventario {
								if item == "Oitão" {
									inventario = slices.Delete(inventario, i, i+1)
									break
								}
							}

							exibirMapa(mapa)
							continue
						} else {
							fmt.Println("\nVocê decide não usar o Oitão. O Predador Alfa te devora! GAME OVER!")
							jogando = false
						}
					} else {
						fmt.Println("\nO Predador Alfa te devora em um único bocado! GAME OVER!")
						jogando = false
					}
				}
			} else if oponenteEncontrado == "Árvore de maças envenenadas" {
				fmt.Println("\nVocê encontrou uma árvore com maçãs deliciosas e comeu uma. Infelizmente, eram envenenadas! GAME OVER!")
				jogando = false
			} else if oponenteEncontrado == "O Pinscher" {
				fmt.Println("\nUm pequeno cachorro aparece latindo furiosamente. É apenas um pinscher, que mal pode fazer?")
				fmt.Println("Você subestima o cachorrinho. Tolo...")
				fmt.Println("Ele começa a latir descontroladamente. Seus latidos ensurdecedores fazem seus tímpanos explodirem enquanto ele devora suas pernas!")
				fmt.Println("GAME OVER! O Pinscher te devorou. Na sua lápide estará escrito: 'Morreu para um cachorro de 2kg'")
				jogando = false
			} else if oponenteEncontrado == "10 mil abelhas perigosas!" {
				fmt.Println("\nVocê ouve um zumbido que rapidamente se torna ensurdecedor.")
				fmt.Println("Uma nuvem negra de abelhas furiosas aparece no horizonte e avança na sua direção!")
				fmt.Println("Em segundos, milhares de ferrões penetram sua pele de todos os lados.")
				fmt.Println("A última coisa que você vê antes de desmaiar é seu corpo transformado em um cacto humano.")
				fmt.Println("GAME OVER! Você virou um cacto de tanto ferrão espetado. Se alguém encontrar seu corpo,")
				fmt.Println("vai achar que é uma nova espécie de planta do deserto.")
				jogando = false
			} else {
				fmt.Printf("\nOh não! Você encontrou %s! GAME OVER!\n", oponenteEncontrado)
				jogando = false
			}
		case "I":
			// Encontrou um item mágico
			itemIndex := rand.Intn(len(itensColocados))
			itemEncontrado := itensColocados[itemIndex]

			// Troca a posição para vazia após coletar o item
			posicaoOriginal = "~~"

			// Atualiza a posição do jogador no mapa
			mapa[jogadorX][jogadorY] = "P"

			// Primeiro exibe o mapa
			exibirMapa(mapa)

			// Depois exibe a mensagem
			fmt.Printf("\nVocê encontrou um item mágico: %s\n", itemEncontrado)
			inventario = append(inventario, itemEncontrado)

		default:
			// 2% de chance do jogador tropeçar e morrer aleatoriamente
			if rand.Intn(100) < 2 {
				fmt.Println("\nVocê move-se confiante para uma nova posição, mas...")
				fmt.Println("TROPEÇA EM UMA PEDRA MINÚSCULA!")
				fmt.Println("Inacreditavelmente, você cai de cabeça e bate na única pedra pontiaguda em um raio de quilômetros.")
				fmt.Println("GAME OVER! Na sua lápide estará escrito: 'Aqui jaz um azarado. Sobreviveu a monstros e perigos, mas foi derrotado pela gravidade'")
				jogando = false
			} else {
				// Atualiza a posição do jogador no mapa
				mapa[jogadorX][jogadorY] = "P"

				// Primeiro exibe o mapa
				exibirMapa(mapa)

				// Depois exibe a mensagem
				fmt.Println("Você se move para uma nova posição...")
			}
		}

		// Eventos aleatórios a cada 5 turnos
		if turnos%5 == 0 {
			eventosAleatorios := []string{
				"Um pássaro pousa na sua cabeça e te confunde com uma árvore por 1 turno",
				"Você pisa em uma banana e escorrega para uma posição aleatória do mapa",
				"Você encontra um espelho mágico e fica admirando seu cabelo perfeito por 1 turno",
				"Uma sereia aparece e te convida para um chá. Depois de recusar, ela fica emburrada e vai embora",
				"Você achou sinal de internet, suficiente pra carregar os reels do instagram. Você assiste um pouco e percebe que ficou uma hora aqui, que loucura, né?",
			}
			fmt.Println(eventosAleatorios[rand.Intn(len(eventosAleatorios))])
		}
	}

	// Revela o mapa completo no final do jogo
	fmt.Println("\nMapa completo:")
	for i := range mapa {
		for j := range mapa[i] {
			if mapa[i][j] == "~~" {
				fmt.Print("~~ ")
			} else {
				fmt.Print(mapa[i][j], " ")
			}
		}
		fmt.Println()
	}

	fmt.Println("\nFim de jogo! Obrigado por jogar!")
}

// Função para exibir o mapa atual com legenda
func exibirMapa(mapa [][]string) {
	fmt.Println("\n  0 1 2 3 4") // Coordenadas horizontais
	for i := range mapa {
		fmt.Print(i, " ") // Coordenadas verticais
		for j := range mapa[i] {
			if mapa[i][j] == "T" || mapa[i][j] == "O" || mapa[i][j] == "I" {
				fmt.Print("~~ ") // Esconde o tesouro, oponentes e itens
			} else {
				fmt.Print(mapa[i][j], " ")
			}
		}
		fmt.Println()
	}

	// Legenda do mapa que aparece em cada atualização
	fmt.Println("\nLegenda:")
	fmt.Println("P = Você | ~~ = Área desconhecida")
	fmt.Println("Os tesouros, oponentes e itens estão escondidos!")
	fmt.Println()
}
