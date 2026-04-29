package http

import (
	"backend/internal/usecase"
	"encoding/json"
	"net/http"
)

type AllocateHandler struct {
	Usecase *usecase.Allocator
}

func (h *AllocateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Orders []any `json:"orders"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	// call usecase
	result := h.Usecase.Auto(nil)

	json.NewEncoder(w).Encode(result)
}