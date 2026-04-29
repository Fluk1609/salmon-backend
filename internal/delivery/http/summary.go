package http

import (
	"backend/internal/usecase"
	"encoding/json"
	"net/http"
)

func SummaryHandler(u *usecase.SummaryUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var orders []any
		json.NewDecoder(r.Body).Decode(&orders)

		res := u.Execute(nil)

		json.NewEncoder(w).Encode(res)
	}
}