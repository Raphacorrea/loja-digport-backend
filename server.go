package main

import (
	//translate to json

	"github.com/Raphacorrea/loja-digport-backend/routes"
)

func StartServer() {

	routes.HandleRequests()
}

/*http.HandleFunc("/produtos", produtosHandler)
	http.ListenAndServe(":8080", nil)}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" { //se nao for vazio-
		produtosFiltradosPorNome := produtosPorNome(queryNome)
		json.NewEncoder(w).Encode(produtosFiltradosPorNome)
	} else if{
		produtos := criaEstoque
		json.NewEncoder(w).Encode(produtos)
	} else if{
		r.Method == "POST"{
		var produto model.Produto
		json.NewDecoder(r.Body).Decode(&produto)
        adicionaProduto(produto)
	}
}
	produtos := criaEstoque()

	json.NewEncoder(w).Encode(produtos)
}

produtos?nome=calca pra filtrar*/
