package main

import (
	"fmt"

	"github.com/Raphacorrea/loja-digport-backend/model"
)

func main() {
	fmt.Println("Hello, world!")

	ListaDeDesejos := model.Carrinho{
		ID:        "ID do carrinho",
		UserID:    "ID do Usuario",
		InfosProduto: []model.InfosProduto{
		
		ProdutoID: "produto1",
		QuantidadeDeProduto: 3,

	},
	{
		ProdutoID: "Produto2",
		QuantidadeDeProduto: 1,
	},
	//fmt.Printf(`%+v`, ListaDeDesejos)
}
//InfosProduto: map[string]int{
//"id-blusa-roxa:1,"
//""id-blusa-rosa:2
//}