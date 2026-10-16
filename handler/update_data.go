package handler

import (
	"encoding/json"
	"net/http"

	"github.com/zikadaz/encurtador/database"
)

// responsável por atualizar um atualizar a url de um código específico
// a implementar

// atualiza a url de um dos códgios

func UpdateCode(w http.ResponseWriter, r *http.Request) {
    key := r.PathValue("code")
	// pega o corpo da requisicao
    var body struct {
        URL string `json:"url"`
    }
	// transforma para formato json
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }
	// Tenta atualizar
    if err := database.UpdateInfoFromDatabase(key, body.URL); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]bool{
            "encontrado": false,
        })
        return
    }
	// caso atualiza, retorna um json
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]bool{
        "atualizado": true,
    })
}