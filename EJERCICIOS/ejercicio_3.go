package main

import (
	"fmt"
	"sync"
)

var cola sync.WaitGroup

func rutina(str string) {
	for i := 0; i < 3; i++ {
		go fmt.Println(str, i)
	}
}

func main() {

	go rutina("rutina 1: ")
	go rutina("rutina 2: ")
	go rutina("rutina 3: ")

	var wait string
	fmt.Scanln(&wait)
}
