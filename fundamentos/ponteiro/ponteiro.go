package main

import "fmt"

func main() {

	i := 1

	var p *int = nil
	p = &i // pegando o endereco da variavel

	fmt.Println(p, *p, i, &i)

	*p++ // desreferenciando (paegando o valor)
	i++

	// Go nao tem aritimetica de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
