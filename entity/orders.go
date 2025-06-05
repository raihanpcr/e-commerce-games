package entity

type OrderDetail struct {
	OrderID     int
	OrderDate   string
	FullName    string
	ProductName string
	Quantity    int
	TotalAmount int
}