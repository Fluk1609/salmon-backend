package domain

type SubOrder struct {
	OrderID     string
	SubOrderID  string
	ItemID      string
	WarehouseID string
	SupplierID  string
	RequestQty  float64
	Type        OrderType
	CreateDate  string
	CustomerID  string
	Allocated   float64
}

type Warehouse struct {
	ID    string
	Name  string
	Stock float64
}

type Customer struct {
	ID          string
	Name        string
	CreditLimit float64
	UsedCredit  float64
}