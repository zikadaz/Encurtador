package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zikadaz/encurtador/database"
)

// responsável por ler models.Urls
// GET /codes


func ReadDataFromDatabase(w http.ResponseWriter, r *http.Request) {
	dataInfo, err := database.ReadInfoFromDatabase("")
	if err != nil {
		fmt.Println("Não foi possível ler o banco de dados.")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataInfo)
}