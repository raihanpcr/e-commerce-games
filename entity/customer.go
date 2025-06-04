package entity

type Customer struct {
	ID      int
	Name    string
	Phone   string
	Address string
	UserID  int
	User    *User
}