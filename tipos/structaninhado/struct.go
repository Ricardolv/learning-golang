package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	toltal := 0.0

	for _, item := range p.itens {
		toltal += item.preco * float64(item.quantidade)
	}

	return toltal
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.10},
			item{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido e R$ %2.f", pedido.valorTotal())
}
