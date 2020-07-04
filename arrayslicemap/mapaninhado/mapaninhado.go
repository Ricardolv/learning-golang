package main

import "fmt"

func main() {

	funcPorLetra := map[string]map[string]float64{
		"B": {
			"Bernardo": 9456.78,
			"Barbara":  5456.78,
		},
		"L": {
			"Liliane": 15456.78,
			"Lili":    8456.78,
		},
		"R": {
			"Richard": 15456.99,
			"Ricardo": 8456.99,
		},
	}

	delete(funcPorLetra, "R")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}

}
