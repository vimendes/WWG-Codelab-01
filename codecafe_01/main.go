package main

import "fmt"

func main() {
	println("coloque suas resoluções nesse arquivo :D")
}

func Apartamentos() {
	apartments := map[string]string{
		"01A": "SHA3948",
		"02A": "DEF8423",
		"01B": "DEF5678",
		"10B": "AEF5678",
		"05D": "DOE5518",
	}
	printPlacas(apartments)

	delete(apartments, "01B")

	fmt.Println()
	printPlacas(apartments)
}

func printPlacas(ap map[string]string) {
	for apart, placa := range ap {
		fmt.Print("Apartamento: ", apart, " | Placa ", placa, "\n")
	}
}
