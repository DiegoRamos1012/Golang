package main

import ("fmt" "errors")

func FibonacciIterativo (n int) []int { // Os parametros são (n int) e retorna um array de inteiros ([]int
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	sequencia := []int{0, 1}
}