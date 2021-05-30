package main

import (
	"fmt"
	"sync"
)

//Declaramos la cola global para poder acceder a ella desde diferentes metodos
var cola sync.WaitGroup

//Metodo que imprime una secuencia de mensajes n veces
func rutina(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, i)
	}
	//Notifico a la cola que esta rutina termino
	cola.Done()
}

//Metodo principal
func main() {
	//Le indicamos a la cola que debe esperar 3 rutinas
	cola.Add(3)
	//Ejecuta 3 tareas de forma concurrente pero una vez que una de ellas
	//inicia las otras se bloan hasta que esta termine imprmiendo un mensaje
	//desde el punto de contro de cada rutina

	go rutina("rutina 1: ")
	go rutina("rutina 2: ")
	go rutina("rutina 3: ")

	//Bloquea el programa para que este no finalize hasta que el metodo Done()
	//sea llamado 3 veces
	cola.Wait()
}
