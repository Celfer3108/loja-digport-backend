package model

type Carrinho struct {
	ProductID    string
	UserID       string
	InfosProduto map[string]int
	ValorTotal   float64
	Quantidade   int
}

type InfosProduto struct {
	ProductID  string
	Quantidade int
}
