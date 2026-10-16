package main

import (
	"log"
	"net/http"

	"github.com/zikadaz/encurtador/handler"

	"github.com/zikadaz/encurtador/database"
)

func main() {
	// conexao com o banco de dados
	database.Connection()
	// rotas/endpoints
	http.HandleFunc("POST /links", handler.AddData)
	http.HandleFunc("GET /links", handler.ReadDataFromDatabase)
	http.HandleFunc("GET /{code}", handler.RedirectUrl)
	log.Println("Rodando....")
	http.ListenAndServe(":8080",nil)
}
