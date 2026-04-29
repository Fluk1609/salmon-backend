package usecase

import "backend/internal/domain"

type Allocator struct{}

func (a *Allocator) Auto(orders []domain.SubOrder) []domain.SubOrder {

	// ❗ optimize: no map recreate, no slice copy
	for i := range orders {
		orders[i].Allocated = orders[i].RequestQty // demo
	}

	return orders
}