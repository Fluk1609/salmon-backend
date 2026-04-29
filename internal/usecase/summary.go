package usecase

import "backend/internal/domain"

type Summary struct {
	TotalOrders    int
	TotalAllocated float64
	TotalRequested float64
	FillRate       float64
}

type SummaryUsecase struct{}

func (s *SummaryUsecase) Execute(orders []domain.SubOrder) Summary {

	var req, alloc float64

	for _, o := range orders {
		req += o.RequestQty
		alloc += o.Allocated
	}

	fill := 0.0
	if req > 0 {
		fill = (alloc / req) * 100
	}

	return Summary{
		TotalOrders:    len(orders),
		TotalAllocated: alloc,
		TotalRequested: req,
		FillRate:       fill,
	}
}