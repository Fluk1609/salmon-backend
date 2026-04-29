package http

import (
	"backend/internal/repository"
	"encoding/json"
	"net/http"
)

func InitHandler(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		o, w2, c := repo.GetAll()

		json.NewEncoder(w).Encode(map[string]interface{}{
			"orders":     o,
			"warehouses": w2,
			"customers":  c,
		})
	}
}