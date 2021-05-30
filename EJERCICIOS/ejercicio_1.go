package main

import (
	"fmt"
)

//Funcion que imprime un texto pasado por parametro
func print(str string) { fmt.Println(str) }

//Metodo princiapal
func main() {
	//Imprime un grupo de mensajes de forma asincrona
	go print("Hola 1")
	go print("Hola 2")
	go print("Hola 3")
	go print("Hola 4")
	go print("Hola 5")
	go print("Hola 6")

	//Hago un scan para bloaquear el programa y este no finalize hasta que
	//el usuario inserte algun caracter
	var wait string
	fmt.Scanln(&wait)
}
