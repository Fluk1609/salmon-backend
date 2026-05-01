package http

import (
	"backend/internal/usecase"
	"backend/internal/domain"
	"encoding/json"
	"net/http"
)

type AllocateHandler struct {
	Usecase *usecase.Allocator
}

type AllocateRequest struct {
	Orders []domain.SubOrder `json:"orders"`
}

func (h *AllocateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var req AllocateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	result := h.Usecase.Auto(req.Orders)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}