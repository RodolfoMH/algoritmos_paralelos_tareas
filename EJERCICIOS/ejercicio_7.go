package main

import (
	"fmt"
	//"sync"
)

func print(str string) {
	fmt.Println(str)
}

//buscar el mayor en el array de forma secuencial usando el metodo de la burbuja
func main() {

	print("Buscando el mayor en: ")

	array := [10]int{2, 3, 6, 7, 8, 1, 23, 34, 4, 5}

	fmt.Println(array)

	var mayor int

	for i := 0; i < len(array); i++ {

		if array[i] > mayor {
			fmt.Println(array[i], "mayor que ", mayor)
			mayor = array[i]
		}
	}

	fmt.Println("El mayor es = ", mayor)
}
