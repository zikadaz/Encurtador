package handler

import (
	"encoding/json"
	"net/http"

	"github.com/zikadaz/encurtador/database"
)

// responsável por apagr um código /  url do banco de dados
// a implementar

// remove um código do
func RemoveInfoFromDatabase(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	err := database.RemoveInfoFromDatabase(code)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]bool{
			"encontrado":false,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{
		"deletado":true,
	})
}