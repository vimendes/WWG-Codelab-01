package main

import "fmt"

func main() {
	fmt.Println("--- EXERCÍCIO 01 ---")
	exercise1()

	fmt.Println("--- EXERCÍCIO 02 ---")
	exercise2()
}

func exercise1() {
	intArray := [10]int{}
	intArray[0] = 2
	intArray[1] = 4
	intArray[2] = 5
	intArray[3] = 6
	intArray[4] = 10
	intArray[5] = 12
	intArray[6] = 14
	intArray[7] = 15
	intArray[8] = 20
	intArray[9] = 24

	fmt.Print(intArray)
	fmt.Println("")

	for i := 0; i < 10; i++ {
		fmt.Println(intArray[i])
	}

}

func exercise2() {
	valuesToBeSum := []int{10, 20, 30, 40, 50}

	fmt.Print("Somatoria = ", sum(valuesToBeSum))
}

func sum(values []int) int {
	sum := 0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum
}
