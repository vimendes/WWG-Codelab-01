package main

import "fmt"

func main() {
	nameStrings()
}

func nameStrings() {
	nomes := [3]string{"vitoria", "erika", "amanda"}
	for _, nome := range nomes {
		saudacao(nome)
	}

	var ano int
	ano = 2019

	fmt.Printf("Binário: %b\n  Decimal: %d\n Hexa: %X", ano, ano, ano)
}

func saudacao(nome string) {
	fmt.Println("Hello, ", nome)
}
