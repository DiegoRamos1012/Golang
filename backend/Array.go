package main

import "fmt"

func PraticandoArrays() { // FUnção desativada na main.go
	numeros := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array original: ", numeros)

	slice := numeros[1:4]
	fmt.Println("Array após modificação do slice: ", slice)

	slice[0] = 10
	fmt.Println("Array após modificação do slice: ", numeros)

	slice = append(slice, 20, 30)
	fmt.Println("Slice após append: ", slice)
}

func PraticandoArrays2() {
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"} // nome da variável - tamanho do array - tipo dos dados no array - objetos no array
	cars[1] = "Opel"
	fmt.Print(cars)
}