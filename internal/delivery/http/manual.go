package http

import (
	"backend/internal/usecase"
	"encoding/json"
	"net/http"
)

type ManualHandler struct {
	Usecase *usecase.ManualUsecase
}

type ManualRequest struct {
	SubOrderID string `json:"subOrderId"`
	Qty        int    `json:"qty"`
}

func (h *ManualHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var req ManualRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	result := h.Usecase.Set(req.SubOrderID, req.Qty)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}