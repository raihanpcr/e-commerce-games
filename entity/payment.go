package entity

type Payment struct {
	PaymentID     int
	OrderID       int64
	PaymentMethod string
	PaymentDate   string
	PaymentPaid   int
}
