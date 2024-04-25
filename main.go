package main

import (
	"fmt"
)

func main() {
	livroscatalago := catalog()
	fmt.Printf("Catalago de livros da Bookworm store: %+v", livroscatalago)
	StartServer()

	nome := ""
	fmt.Scan("Digite o nome do produto desejado", &nome)
	produtosFiltrados := buscaPorNome(nome)

	fmt.Println(produtosFiltrados)
}
