package usecase

import "backend/internal/domain"

type Summary struct {
	TotalOrders    int     `json:"totalOrders"`
	TotalAllocated int     `json:"totalAllocated"`
	TotalRequested int     `json:"totalRequested"`
	FillRate       float64 `json:"fillRate"`
}

type SummaryUsecase struct{}

func (s *SummaryUsecase) Execute(orders []domain.SubOrder) Summary {

	var req, alloc int

	for _, o := range orders {
		req += o.RequestQty
		alloc += o.Allocated
	}

	fill := 0.0
	if req > 0 {
		fill = (float64(alloc) / float64(req)) * 100
	}

	return Summary{
		TotalOrders:    len(orders),
		TotalAllocated: alloc,
		TotalRequested: req,
		FillRate:       fill,
	}
}
