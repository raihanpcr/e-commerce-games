package entity

type Customer struct {
	ID      int
	Name    string
	Phone   string
	Address string
	UserID  int
	User    *User
}

type MostOrderCustomer struct {
	CustomerID   int
	CustomerName string
	TotalOrders  int
}