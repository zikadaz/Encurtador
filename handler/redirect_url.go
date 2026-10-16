package handler

// responsável por redirecionar com base no código
// GET /{code}
import (	
	"net/http"
	"github.com/zikadaz/encurtador/database"

)

func RedirectUrl(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	result, err := database.ReadInfoFromDatabase(code)
	if err != nil {
		http.Error(w, `{"failed":true}`,
			http.StatusNotFound)
		return
	}

	if err := database.IncrementClicks(code); err != nil {
		http.Error(w, "Erro ao atualizar cliques",
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, result[code].Url, http.StatusFound)
}