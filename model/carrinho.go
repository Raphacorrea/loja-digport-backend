package model

type Carrinho struct {
	ProdutoID    []string
	UserID       string
	ID           string
	InfosProduto []InfosProduto
	//InfosProduto map[string]int
	ValorTotal      float64
	QuantidadeTotal int
}
type InfosProduto struct {
	ProdutoID           string
	QuantidadeDeProduto int
}
