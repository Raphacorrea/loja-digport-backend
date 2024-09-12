package main

import (
	"errors"

	"github.com/Raphacorrea/loja-digport-backend/model"
)

var Estoque []model.Produto = []model.Produto{}

func produtosPorNome(nome string) []model.Produto {
	resultado := []model.Produto{}

	for _, produto := range produtos {
		if produto.Nome == nome {
			resultado = append(resultado, produto)
		}
	}

	return resultado
}
func adicionaProduto(produtoNovo model.Produto) error {
	for _, produto := range produtos {
		if produtoNovo.ID == produto.ID {
			return errors.New("o id é inválido")
		}
	}
	produtos = append(produtos, produtoNovo)
	return nil
}

func criaEstoque() []model.Produto {
	produtos := []model.Produto{
		{
			Nome:       "Calça tradicional",
			Descricao:  "Calça Jeans",
			Categoria:  "calça",
			ID:         "a1",
			Valor:      100.00,
			Quantidade: 1,
			Imagem:     "calça.jpg",
		},
		{
			Nome:       "Blusa verão",
			Descricao:  "Blusa verão manga curta",
			Categoria:  "camisa",
			ID:         "a2",
			Valor:      70.00,
			Quantidade: 1,
			Imagem:     "blusa.jpg",
		},
	}
	return produtos
}
