package main

import (
	"fmt"
	"math"
)

//Funcion que imprime un texto pasado por parametro
func print(str string) { fmt.Println(str) }

//Variable global compartida entre rutinas que almacenara el mayor numero encontrado
var mayor int

//Esta funcion recibe un array y busca desde la izquierda hacia la derecha el mayor numero
//y se detiene en el centro.
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

//Esta funcion recibe un array y busca desde la derecha hacia la izquierda el mayor numero
//y se detiene en el centro.
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

//Metodo principal
func main() {

	//buscar el mayor numero en el arreglo para esto lo divide en 2 partes y comeinza
	//a recorrerlo por izquierda y derecha hasta terminar en la mitad
	print("Buscando el mayor en: ")
	array := [10]int{2, 3, 6, 7, 8, 1, 23, 34, 4, 5}
	fmt.Println(array)

	//busco el mayor recorriendo la parte derecha y izquierda del array al mismo tiempo
	go buscarDer(array)
	go buscarIzq(array)

	//Hago un scan para bloaquear el programa y este no finalize hasta que el usuario
	//inserte algun caracter
	var wait string
	fmt.Scanln(&wait)

	//Imprimo el mayor
	fmt.Println("El mayor es = ", mayor)
}
