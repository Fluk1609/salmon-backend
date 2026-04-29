package http

import (
	"backend/internal/domain"
	"backend/internal/usecase"
	"encoding/json"
	"net/http"
)

// ===== Request / Response DTO =====

type ManualRequest struct {
	SubOrderID string             `json:"subOrderId"`
	Qty        float64            `json:"qty"`
	Orders     []domain.SubOrder  `json:"orders"`
}

type ManualResponse struct {
	OK     bool                `json:"ok"`
	Orders []domain.SubOrder   `json:"orders"`
}

// ===== Handler =====

func ManualHandler(uc *usecase.ManualUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// ✅ method check
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req ManualRequest

		// ✅ decode + validate
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if req.SubOrderID == "" {
			http.Error(w, "subOrderId is required", http.StatusBadRequest)
			return
		}

		if req.Qty < 0 {
			http.Error(w, "qty must be >= 0", http.StatusBadRequest)
			return
		}

		// ✅ call usecase
		updated := uc.Execute(
			req.SubOrderID,
			req.Qty,
			req.Orders,
		)

		// check success (หา order ที่แก้)
		ok := false
		for _, o := range updated {
			if o.SubOrderID == req.SubOrderID {
				ok = true
				break
			}
		}

		// ✅ response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ManualResponse{
			OK:     ok,
			Orders: updated,
		})
	}
}