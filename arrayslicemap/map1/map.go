package main

import "fmt"

func main() {

	//var aprovados map[int] string
	//mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678888] = "Liliane"
	aprovados[12345678999] = "Bernardo"
	aprovados[12345678910] = "Maressa"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678888])
	delete(aprovados, 12345678888)
	fmt.Println(aprovados[12345678888])

}
