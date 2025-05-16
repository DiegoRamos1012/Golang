// Golang program to illustrate the usage of
// Sleep() function

// Including main package
package main

// Importing time and fmt
import (
    "fmt" // Pacote de formatação
    "time" // Pacote de exibição de tempo
)

// Main function
func RelogioCore() {
	for {
		agora := time.Now() // Retorna o horário atual

		fmt.Printf("\r%02d:%02d:%02d", agora.Hour(), agora.Minute(), agora.Second()) // fmt com Printf pode usar formatadores
		// O "\r" edita a mesma linha ao invés de criar uma nova.
		// O "%d" indica que um número inteiro decimal será formatado. Exemplos: 9 -> "09", 12 -> "12", 3 -> "03"
		// O "0" preenche com zeros à esquerda ao invés de espaços
		// O "2" define a larguma mínima como 2 caracteres

		time.Sleep(1 * time.Second)
	}
}

func Relogio() {
	fmt.Println("Relógio Digital CLI muito brabo")
	fmt.Println("Pressione Ctrl + C para sair")
	RelogioCore() // Executa a função declarada anteriormente
}