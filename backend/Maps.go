package main

import "fmt"

func Maps() {
	// Cria um mapa onde a chave e o valor são strings
	capitais := map[string]string{
		"Brasil":   "Brasília",
		"França":   "Paris",
		"Alemanha": "Berlim",
	}

	// Adiciona uma nova chave "Japão" com o valor "Tóquio"
	capitais["Japão"] = "Tóquio"

	// Imprime a capital do Brasil acessando pelo nome do país
	fmt.Println("Capital do Brasil: ", capitais["Brasil"])

	// Percorre o mapa e imprime cada país e sua capital
	for pais, capital := range capitais {
		fmt.Printf("País: %s, Capital: %s\n", pais, capital)
	}
}
