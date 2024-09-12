package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//StartServer()
	//fmt.Printf("Esse é o catálogo da loja listaProdutos:%+v", estoque())
	db := ConectaBancoDados()
	fmt.Println("Bem Vindos à Loja DigPort!")

	defer db.Close()
}

func ConectaBancoDados() *sql.DB {
	connStr := "user=postgres dbname=digport_loja password=digport host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
