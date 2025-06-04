package entity

type OrderProduct struct {
	OrderID     int
	CustomerID  int
	OrderItem   []OrderItem
	TotalAmount int
	OrderDate   string
	Status      string
}

type OrderItem struct {
	OrderItemID int
	OrderID     int
	ProductID   int
	Quantity    int
	UnitPrice   int
	Subtotal    int
}

type Order struct {
	OrderID     int
	CustomerID  int
	OrderDate   string
	Status      string
	TotalAmount int
}
