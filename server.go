package main

import (
	"encoding/json"
	"net/http"

	"github.com/Celfer3108/DigiPort/model"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getProdutos(w, r)
	} else if r.Method == "POST" {
		addProdutos(w, r)
	}
}

func addProdutos(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	registerbooks(produto)

	w.WriteHeader(http.StatusCreated)
}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" {
		booksFiltradosPorNome := buscaPorNome(queryNome)
		json.NewEncoder(w).Encode(booksFiltradosPorNome)
	} else {
		produtos := product
		json.NewEncoder(w).Encode(produtos)
	}
}
