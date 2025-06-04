package handler

import "database/sql"

type CustomerHandler struct {
	DB *sql.DB
}