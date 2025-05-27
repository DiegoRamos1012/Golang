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

func Slice_Array() {
	// Array
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println("Array: ", numbers)

	// Slice
	frutas := []string{"Banana", "Maçã", "Laranja", "Manga"}
	fmt.Println("Frutas: ", frutas)
}

func PraticandoArrays2() {
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"} // variável declarada - nome da variável - tamanho do array - tipo dos dados no array - objetos no array
	cars[1] = "Opel"
	fmt.Print(cars)
}
