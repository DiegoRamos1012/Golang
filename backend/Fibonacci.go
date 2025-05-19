package main

import ("fmt")

func FibonacciIterativo (n int) []int { // Os parametros são (n int) e retorna um array de inteiros ([]int)
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	sequencia := []int{0, 1} // Inicializa a sequência com os dois primeiros números

	for i := 2; i < n; i++ {
		proximo := sequencia[i-1] + sequencia[i-2]
		sequencia = append(sequencia, proximo) // Adiciona o próximo número
	}

	return sequencia
}

func executarFibonacci() {
	n := 10 // Quantidade de números na sequência
	fmt.Println("Fibonacci (Iterativo): ", FibonacciIterativo(n))
}