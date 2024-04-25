package main

import (
	"errors"

	"github.com/Celfer3108/DigiPort/model"
)

var product []model.Produto

func catalog() []model.Produto {
	livros := []model.Produto{
		{
			Nome:       "A paciente silenciosa",
			Autor:      "Alex Michaelides",
			Genero:     "Thriller",
			Editora:    "Record",
			ID:         0001,
			Quantidade: 3,
			Valor:      49.90,
			Imagem:     "livro.jpg",
		},

		{
			Nome:       "Daisy Jones and the six",
			Autor:      "Alex Michaelides",
			Genero:     "Drama",
			Editora:    "aralela",
			ID:         0002,
			Quantidade: 5,
			Valor:      39.90,
			Imagem:     "livro.jpg",
		},
		{
			Nome:       "O iluminado",
			Autor:      "Stephen King",
			Genero:     "Suspense",
			Editora:    "Suma",
			ID:         0003,
			Quantidade: 2,
			Valor:      69.90,
			Imagem:     "livro.jpg",
		},
		{
			Nome:       "A sombra do vento",
			Autor:      "Carlos Ruiz Zafon",
			Genero:     "Romance",
			Editora:    "Suma",
			ID:         0004,
			Quantidade: 5,
			Valor:      49.90,
			Imagem:     "livro.jpg",
		},
		{
			Nome:       "Ken Follet",
			Autor:      "Nunca",
			Genero:     "Drama",
			Editora:    "Arqueiro",
			ID:         0005,
			Quantidade: 3,
			Valor:      79.90,
			Imagem:     "livro.jpg",
		},
	}
	return livros

}

func buscaPorNome(nome string) []model.Produto {
	resultado := []model.Produto{}

	for _, produto := range product {
		if produto.Nome == nome {
			resultado = append(resultado, produto)
		}
	}
	return resultado
}
func registerbooks(books model.Produto) error {
	for _, produto := range product {
		if books.ID == produto.ID {
			return errors.New("")
		}
	}
	product = append(product, books)
	return nil
}
