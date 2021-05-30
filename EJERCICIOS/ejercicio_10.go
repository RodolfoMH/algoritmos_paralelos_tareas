package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//Funcion que imprime un texto pasado por parametro
func print(str string) {
	fmt.Println(str)
}

//Declaramos la cola global para poder acceder a ella desde diferentes metodos
var cola sync.WaitGroup

//Funcion que genera un numero al azar dentro del rango especificado.
func generarRandon(posicion string, rango int) {
	fmt.Println(posicion, rand.Intn(rango))
	cola.Done()
}

//Metodo principal
func main() {

	//generar cada numero del loto de forma asincrona usando numeros randon

	fmt.Println("Los 7 numeros ganadores son: ")
	//Le indicamos a la cola que debe esperar 7 rutinas
	cola.Add(7)

	go generarRandon("Primer numero => ", 38)
	go generarRandon("segundo numero => ", 38)
	go generarRandon("tercero numero => ", 38)
	go generarRandon("cuarto numero => ", 38)
	go generarRandon("quinto numero => ", 38)
	go generarRandon("sexto numero => ", 38)

	go generarRandon("loto mas => ", 10)

	//Bloquea el programa para que este no finalize hasta que el metodo Done() sea llamado 7 veces
	cola.Wait()

}
