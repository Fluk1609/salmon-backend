package domain

type SubOrder struct {
	OrderID     string
	SubOrderID  string
	ItemID      string
	WarehouseID string
	SupplierID  string
	RequestQty  int
	Type        OrderType
	CreateDate  string
	CustomerID  string
	Allocated   int
}

type Warehouse struct {
	ID    string
	Name  string
	Stock int
}

type Customer struct {
	ID          string
	Name        string
	CreditLimit int
	UsedCredit  int
}