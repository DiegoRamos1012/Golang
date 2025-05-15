package main

import (
	"fmt"
)

func CelsiusParaFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func FahrenheitParaCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func CelsiusParaKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func KelvinParaCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func ConversorGraus() {
	var inputNum, escolha string
	var num float64

	fmt.Println("\n==== Conversor de temperatura ====")
	fmt.Println("1 - Celsius para Fahrenheit")
	fmt.Println("2 - Fahrenheit para Celsius")
	fmt.Println("3 - Celsius para Kelvin")
	fmt.Println("4 - Kelvin para Celsius")
	fmt.Print("Escolha uma opção (1-4): ")
	fmt.Scanln(&escolha)

	fmt.Print("Digite a temperatura: ")
	fmt.Scanln(&inputNum)

	// Converter string para float64
	fmt.Sscanf(inputNum, "%f", &num)

	switch escolha {
	case "1":
		resultado := CelsiusParaFahrenheit(num)
		fmt.Printf("%.2f°C = %.2f°F\n", num, resultado)
	case "2":
		resultado := FahrenheitParaCelsius(num)
		fmt.Printf("%.2f°F = %.2f°C\n", num, resultado)
	case "3":
		resultado := CelsiusParaKelvin(num)
		fmt.Printf("%.2f°C = %.2f K\n", num, resultado)
	case "4":
		resultado := KelvinParaCelsius(num)
		fmt.Printf("%.2f K = %.2f°C\n", num, resultado)
	default:
		fmt.Println("Opção inválida!")
	}
}
