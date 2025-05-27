package main

import (
	"fmt"
	"sync"
	"time"
)

func dizer(texto string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for range 3 {
		fmt.Println(texto)
		time.Sleep(time.Millisecond * 500) // Pausa de 500ms
	}
}

func Goroutines() {
	var wg sync.WaitGroup
	wg.Add(1)
	go dizer("Olá", &wg) // Goroutine
	dizer("Mundo", nil)  // Executa de forma síncrona
	wg.Wait()
}
