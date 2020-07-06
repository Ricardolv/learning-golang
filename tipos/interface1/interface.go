package main

import "fmt"

type impremivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobreNome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces sao implementada implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobreNome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x impremivel) {
	fmt.Println(x.toString())
}

func main() {

	var coisa impremivel = pessoa{"Ricardo", "Viana"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Laptop Avell", 11000.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Laptop Avell", 7000.90}
	imprimir(p2)

}
