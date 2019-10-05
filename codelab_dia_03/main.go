package main

import "fmt"

func main() {
	println("coloque suas resoluções nesse arquivo :D")
}

func NumberArray() {
	var numeros = [10]string{"zero", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove"}
	for i, num := range numeros {
		fmt.Println(i, " - ", num)
	}
}

func ListaCompras() {

	var lista []string
	var lista2 = []string{"Alho", "Café", "Cebola"}

	lista = append(lista, lista2...)
	lista = append(lista, "Chocolate", "Feijão", "Frango")

	for _, coisa := range lista {
		fmt.Println(coisa)
	}

	fmt.Println("------------")

	var listaParte = lista[2:4]
	for _, coisa := range listaParte {
		fmt.Println(coisa)
	}
}

func SomatoriaSlices() {
	var slice1 = []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	var slice2 = []int{1, 2, 4, 6, 8, 10, 12, 14, 16, 18}

	var sliceSoma []int
	var sliceTotal []int
	var i int
	for i = 0; i < 10; i = i + 1 {
		sliceSoma = append(sliceSoma, slice1[i]+slice2[i])
	}

	sliceTotal = append(sliceTotal, slice1...)
	sliceTotal = append(sliceTotal, slice2...)

	var sliceResultado1 = sliceSoma[5:12]
	var sliceResultado2 = sliceTotal[5:13]

	fmt.Println(sliceResultado1)
	fmt.Print(sliceResultado2)
}
