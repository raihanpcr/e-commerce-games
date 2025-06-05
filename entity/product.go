package entity

type Product struct {
	ProductID   int
	Name        string
	Description string
	Price       int
	Stock       int
}

type BestSellingProduct struct {
	ProductID   int
	ProductName string
	Terjual     int
}

type MustBeRestockProduct struct {
	ProductID int
	Name      string
	Stock     int
}
