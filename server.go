package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8080", nil)

}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" { //se nao for vazio-
		produtosFiltradosPorNome := produtosPorNome(queryNome)
		json.NewEncoder(w).Encode(produtosFiltradosPorNome)
	} else {
		produtos := ListaProdutos
		json.NewEncoder(w).Encode(produtos)
	}
	produtos := criaEstoque()

	json.NewEncoder(w).Encode(produtos)
}

//produtos?nome=calca pra filtrar
