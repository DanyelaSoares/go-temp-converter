package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-temp-converter/converter"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=== Conversor de Temperatura ===")
		fmt.Println("1 - Celsius para Fahrenheit")
		fmt.Println("2 - Fahrenheit para Celsius")
		fmt.Println("3 - Celsius para Kelvin")
		fmt.Println("4 - Kelvin para Celsius")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		optionStr, _ := reader.ReadString('\n')
		optionStr = strings.TrimSpace(optionStr)
		option, _ := strconv.Atoi(optionStr)

		if option == 0 {
			fmt.Println("Saindo...")
			break
		}

		fmt.Print("Digite a temperatura: ")
		valueStr, _ := reader.ReadString('\n')
		valueStr = strings.TrimSpace(valueStr)
		value, _ := strconv.ParseFloat(valueStr, 64)

		switch option {
		case 1:
			fmt.Printf("Resultado: %.2f ºF\n", converter.CelsiusToFahrenheit(value))
		case 2:
			fmt.Printf("Resultado: %.2f ºC\n", converter.FahrenheitToCelsius(value))
		case 3:
			fmt.Printf("Resultado: %.2f K\n", converter.CelsiusToKelvin(value))
		case 4:
			fmt.Printf("Resultado: %.2f ºC\n", converter.KelvinToCelsius(value))
		default:
			fmt.Println("Opção inválida.")
		}

		fmt.Println()
	}
}
