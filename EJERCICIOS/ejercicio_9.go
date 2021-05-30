package main

import (
	"fmt"
	//"sync"
)

func print(str string) {
	fmt.Println(str)
}

func main() {

	//sumar un numero a cada elemento de un array

	array1 := [10]int{2, 3, 6, 7, 8, 1, 23, 34, 4, 5}
	array2 := [10]int{5, 4, 34, 23, 1, 8, 7, 6, 3, 2}

	fmt.Println("Array 1", array1)
	fmt.Println("Array 2", array2)

	arrays1 := [10]int{}
	arrays2 := [10]int{}
	arrays3 := [10]int{}

	for i := 0; i < len(array1); i++ {
		fmt.Println("Sumando ", array1[i], "con", array2[i], " = ", array1[i]+array2[i])
		fmt.Println("Resta ", array2[i], "con", array1[i], " = ", array2[i]+array1[i])
		fmt.Println("Sumando2 ", array1[i], "con", array1[i], " = ", array1[i]+array1[i])
		arrays1[i] = array1[i] + array2[i]
		arrays2[i] = array2[i] - array1[i]
		arrays3[i] = array1[i] + array1[i]
	}
	fmt.Println("Resultado Suma", arrays1)
	fmt.Println("Resultado Resta", arrays2)
	fmt.Println("Resultado Suma2", arrays3)

}
