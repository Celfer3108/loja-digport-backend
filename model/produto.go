package model

type Produto struct {
	Nome       string  `json:"Nome"`
	Autor      string  `json:"Autor"`
	Genero     string  `json:"Genero"`
	Editora    string  `json:"Editora"`
	ID         string  `json:"ID"`
	Quantidade int     `json:"Quantidade"`
	Valor      float64 `json:"Valor"`
	Imagem     string  `json:"Imagem"`
}
