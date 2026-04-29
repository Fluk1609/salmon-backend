package dto

import "backend/internal/usecase"

type SummaryResponse struct {
	Data usecase.Summary `json:"data"`
}