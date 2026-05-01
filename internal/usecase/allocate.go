package usecase

import "backend/internal/domain"

type Allocator struct{}

type AllocateResult struct {
	Orders     []domain.SubOrder `json:"orders"`
	Warehouses []any             `json:"warehouses"`
	Customers  []any             `json:"customers"`
	Logs       []any             `json:"logs"`
}

func (a *Allocator) Auto(orders []domain.SubOrder) AllocateResult {

	for i := range orders {
		orders[i].Allocated = orders[i].RequestQty
	}

	return AllocateResult{
		Orders:     orders,
		Warehouses: []any{},
		Customers:  []any{},
		Logs:       []any{},
	}
}