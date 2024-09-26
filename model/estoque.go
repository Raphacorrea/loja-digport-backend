package model

<<<<<<< HEAD
import "github.com/grasieliw/loja-digport-backend/model"

//import "github.com/Raphacorrea/loja-digport-backend/model"

=======
//import "github.com/Raphacorrea/loja-digport-backend/model"

>>>>>>> 207ffc60e944534e18651511c17efab8c3530b40
func criaEstoque() {
	produtos := []model.Produto{
		{
			ID:                  "a1",
			Nome:                "Calça tradicional",
			Preco:               100.00,
			Descricao:           "Calça Jeans Azul Escuro",
			Categoria:           "Calça",
			Imagem:              "calça.jpg",
			QuantidadeEmEstoque: 5,
		},
		{
			ID:                  "a2",
			Nome:                "Blusa verão",
			Preço:               70.00,
			Descricao:           "Blusa verão manga curta",
			Categoria:           "camisa",
			Imagem:              "blusa.jpg",
			QuantidadeEmEstoque: 1,
		},
	}

	/*
		type Estoque map[string]model.Produto
		qtdInicialProdutos :=
		estoque := make(map[string]model.Produto, qtdInicialProdutos)
		estoque := make(Estoque, qtdInicialProdutos)
		estoque2["1"] = model.Produto{ID: "1", Nome: "Revista Recreio", Preco: 19.90, Descricao: "Revista Recreio", Imagem: "revista.jpg"}
	*/

	estoque := make([]model.Produto, len(produtos))

	return estoque
}
