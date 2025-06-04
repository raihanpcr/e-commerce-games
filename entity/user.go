package entity

import "database/sql"

type User struct {
	ID       int
	Email    string
	Password string
	Role     string
	Token    sql.NullString
}