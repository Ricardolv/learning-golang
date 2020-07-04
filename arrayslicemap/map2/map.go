package main

import "fmt"

func main() {

	funcsESalarios := map[string]float64{
		"Liliane":  11325.45,
		"Bernardo": 15456.78,
		"Maressa":  1200.0,
	}

	fmt.Println(funcsESalarios)

	funcsESalarios["Ricardo"] = 1350.0

	fmt.Println(funcsESalarios)

	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
