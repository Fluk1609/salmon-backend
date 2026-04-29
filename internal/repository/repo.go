package repository

import "backend/internal/domain"

type Repository interface {
	GetAll() ([]domain.SubOrder, []domain.Warehouse, []domain.Customer)
	Reset() ([]domain.SubOrder, []domain.Warehouse, []domain.Customer)
}