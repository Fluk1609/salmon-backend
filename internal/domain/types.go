package domain

type OrderType string

const (
	EMERGENCY OrderType = "EMERGENCY"
	OVERDUE   OrderType = "OVERDUE"
	DAILY     OrderType = "DAILY"
)