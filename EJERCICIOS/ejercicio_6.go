package main

import (
	"fmt"
	"math"
	//"sync"
)

func print(str string) {
	fmt.Println(str)
}

//var cola sync.WaitGroup

var mayor int

func buscarIzq(array [10]int) {

	print("Buscando de izquierda a derecha")

	var hasta int

	hasta = int(math.Round(float64(len(array) / 2)))

	for i := 0; i < hasta; i++ {

		if array[i] > mayor {
			fmt.Println("Desde la Izquierda  =>  ", array[i], "mayor que ", mayor)
			mayor = array[i]
		}

	}
}

func buscarDer(array [10]int) {

	print("Buscando de derecha a izquierda")

	var hasta int

	hasta = int(math.Round(float64(len(array) / 2)))

	for i := (len(array) - 1); i > hasta; i-- {

		if array[i] > mayor {
			fmt.Println("Desde la Derecha  =>  ", array[i], "mayor que ", mayor)
			mayor = array[i]
		}

	}
}

//buscar el mayor numero en el arreglo para esto lo divide en 2 partes y comeinza a recorrerlo por izquierda y derecha hasta terminar en la mitad
func main() {

	print("Buscando el mayor en: ")

	array := [10]int{2, 3, 6, 7, 8, 1, 23, 34, 4, 5}

	fmt.Println(array)

	go buscarIzq(array)
	go buscarDer(array)

	var wait string
	fmt.Scanln(&wait)

	fmt.Println("El mayor es = ", mayor)
}
