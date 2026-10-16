package handler
// responsável por criar um novo link
// POST /code {url:"https:..."}
import (
	"encoding/json"
	"net/http"

	"github.com/zikadaz/encurtador/model"
	"github.com/zikadaz/encurtador/utils"
	"github.com/zikadaz/encurtador/database"
)


func AddData(w http.ResponseWriter, r *http.Request) {
	// verifica se a chamada foi realmente um post
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}
	var m model.Url
	// decodifica o texto json com base no modelo
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Erro ao ler os dados enviados: "+err.Error(), http.StatusBadRequest)
		return
	}
	if m.Url == "" {
		http.Error(w, "A url é obrigatória.", http.StatusBadRequest)
		return
	}
	m.Code = utils.GeradorAleatorio()
	database.AddInfoToDatabase(m.Code, m.Url)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}