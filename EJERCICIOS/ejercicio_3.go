package main

import (
	"fmt"
)

//Metodo que imprime una secuencia de mensajes n veces de forma asincrona
func rutina(str string) {
	for i := 0; i < 3; i++ {
		go fmt.Println(str, i)
	}
}

func main() {

	//se imprime un secuencias de mensajes de contros desde diferentes rutinas de manera simultanea
	go rutina("rutina 1: ")
	go rutina("rutina 2: ")
	go rutina("rutina 3: ")

	//Hago un scan para bloaquear el programa y este no finalize hasta que el usuario inserte algun caracter
	var wait string
	fmt.Scanln(&wait)
}
