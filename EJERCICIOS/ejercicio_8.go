package main

import (
	"fmt"
	//"sync"
)

func print(str string) {
	fmt.Println(str)
}

func sumarArray(array1 [10]int, array2 [10]int) {

	array := [10]int{}
	for i := 0; i < len(array1); i++ {
		fmt.Println("Sumando ", array1[i], "con", array2[i], " = ", array1[i]+array2[i])
		array[i] = array1[i] + array2[i]
	}
	fmt.Println("Resultado Suma", array)
}

func restaArray(array1 [10]int, array2 [10]int) {

	array := [10]int{}
	for i := 0; i < len(array1); i++ {
		fmt.Println("Restando ", array1[i], "con", array2[i], " = ", array1[i]-array2[i])
		array[i] = array1[i] - array2[i]
	}
	fmt.Println("Resultado Resta", array)
}

//Sumna o resta 2 o mas arrays de forma simultanea
func main() {

	array1 := [10]int{2, 3, 6, 7, 8, 1, 23, 34, 4, 5}
	array2 := [10]int{5, 4, 34, 23, 1, 8, 7, 6, 3, 2}

	fmt.Println("Array 1", array1)
	fmt.Println("Array 2", array2)

	go sumarArray(array1, array2)
	go restaArray(array2, array1)
	go sumarArray(array1, array1)

	var wait string
	fmt.Scanln(&wait)
}
