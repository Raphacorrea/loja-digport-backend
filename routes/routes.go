package routes

import (
	"net/http"

	"github.com/Raphacorrea/loja-digport-backend/controller"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRequests() {

	route := mux.NewRouter()

	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.BuscaProdutosPorNomeHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.CriaProdutosHandler).Methods("POST")
	route.HandleFunc("/produtos", controller.RemoveProdutoHandler).Methods("DELETE")
	//route.HandleFunc("/produtos", controller.AtualizaProdutoHandler).Methods("PUT")
	//uusuario
	route.HandleFunc("/usuarios", controller.CriaUsuarioHandler).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(route)
	http.ListenAndServe(":8080", handler)
	//localhost:8080
}
