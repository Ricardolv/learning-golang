package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// numericos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro e", reflect.TypeOf(32000))

	// sem sinal (so positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte e", reflect.TypeOf(b))

	// com sinal (so positivos)... uint8 uint16 uint32 uint64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int e", i1)
	fmt.Println("O tipo de i1 e", reflect.TypeOf(i1))

	// representa um mapeamento da tabela Unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O rune e", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x e", reflect.TypeOf(x))
	fmt.Println("O tipo do literal de x e", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo e", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Opa, meu nome e Bernardo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string e", len(s1))

	// string com multiplas linhas
	s2 := `Ola
	meu
	nome
	e
	Bernardo`
	fmt.Println(s2 + "!")

	// char????
	char := 'a'
	fmt.Println("O tamanho do char e", reflect.TypeOf(char))
	fmt.Println(char)

}
