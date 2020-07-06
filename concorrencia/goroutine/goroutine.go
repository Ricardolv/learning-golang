package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {

	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (itercao %d)\n", pessoa, texto, i+1)
	}
}

func main() {

	// fale("Liliane", "pq vc nao fala comigo?", 3)
	// fale("Ricardo", "So posso falar depois de vc!", 1)

	// go fale("Liliane", "Ei...", 500)
	// go fale("Ricardo", "Opa...", 500)

	go fale("Liliane", "Entendi!!!", 10)
	fale("Ricardo", "Parabens!", 5)

}
