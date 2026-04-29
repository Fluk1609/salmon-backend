package repository

import (
	"backend/internal/domain"
	"fmt"
)

type MemoryRepo struct{}

func (m *MemoryRepo) GetAll() ([]domain.SubOrder, []domain.Warehouse, []domain.Customer) {

	wh := []domain.Warehouse{
		{ID: "WH-001", Name: "BKK", Stock: 500},
		{ID: "WH-002", Name: "CNX", Stock: 300},
	}

	cust := []domain.Customer{
		{ID: "CT-1", Name: "A", CreditLimit: 1000},
		{ID: "CT-2", Name: "B", CreditLimit: 2000},
	}

	var orders []domain.SubOrder

	for i := 1; i <= 5000; i++ {
		id := fmt.Sprintf("ORDER-%04d", i)

		orders = append(orders, domain.SubOrder{
			OrderID:     id,
			SubOrderID:  id + "-001",
			RequestQty:  float64(i % 100),
			Type:        domain.DAILY,
			CustomerID:  cust[i%2].ID,
			WarehouseID: wh[i%2].ID,
		})
	}

	return orders, wh, cust
}

func (m *MemoryRepo) Reset() ([]domain.SubOrder, []domain.Warehouse, []domain.Customer) {
	return m.GetAll()
}