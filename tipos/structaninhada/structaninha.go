package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() (total float64) {
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, qtde: 2, preco: 12.10},
			{produtoID: 2, qtde: 1, preco: 23.49},
			{produtoID: 11, qtde: 100, preco: 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é: R$%.2f", pedido.valorTotal())
}
