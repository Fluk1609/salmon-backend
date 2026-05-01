package usecase

import "backend/internal/domain"

type ManualUsecase struct{}

type ManualResult struct {
	Orders []domain.SubOrder `json:"orders"`
}

func (u *ManualUsecase) Set(subOrderId string, qty int) ManualResult {

	orders := []domain.SubOrder{}

	for i := range orders {
		if orders[i].SubOrderID == subOrderId {
			orders[i].Allocated = qty
		}
	}

	return ManualResult{
		Orders: orders,
	}
}