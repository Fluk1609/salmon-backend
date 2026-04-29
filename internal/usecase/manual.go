package usecase

import "backend/internal/domain"

type ManualUsecase struct{}

func (u *ManualUsecase) Execute(
	id string,
	qty float64,
	orders []domain.SubOrder,
) []domain.SubOrder {

	for i := range orders {
		if orders[i].SubOrderID == id {
			orders[i].Allocated = qty
		}
	}
	return orders
}