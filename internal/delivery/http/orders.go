package http

import (
	"backend/internal/repository"
	"encoding/json"
	"net/http"
)

func OrdersHandler(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		o, _, _ := repo.GetAll()

		json.NewEncoder(w).Encode(o[:50])
	}
}