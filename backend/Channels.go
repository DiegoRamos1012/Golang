package main // Define o pacote principal do programa

import "fmt" // Importa o pacote fmt para impressão no console

// Função que soma dois inteiros e envia o resultado para um canal
func soma(a int, b int, ch chan int) {
	resultado := a + b // Soma os valores recebidos
	ch <- resultado    // Envia o resultado para o canal
}

func Channels() {
	canal := make(chan int) // Cria um canal para transmitir valores do tipo int

	go soma(3, 4, canal) // Executa a função soma em uma goroutine, passando o canal

	resultado := <-canal                 // Recebe o valor enviado pela goroutine através do canal
	fmt.Println("Resultado:", resultado) // Imprime o resultado recebido
}
