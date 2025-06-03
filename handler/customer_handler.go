package handler

import (
	"database/sql"
	"log"
	"strings"
)

type CustomerHandler struct {
	DB *sql.DB
}

func (h *CustomerHandler) AddCustomer(email, password, fullname, phonenumber, address string){

	role := "Customer"
	//Simpan data
	query, err := h.DB.Exec(`
		INSERT INTO users (email, password, role) VALUES (?,?,?)`,
		strings.TrimSpace(email),
		strings.TrimSpace(password),
		strings.TrimSpace(role),
	)

	if err != nil {
		log.Println("Gagal Add Users : ", err)
	}else{
		log.Println("User Berhasil Ditambahkan")
	}

	userID, err := query.LastInsertId()
	if err != nil {
		log.Fatal("Gagal ngambil user_id %w",err)
	}

	_, err = h.DB.Exec(`INSERT INTO customers (user_id, full_name, phone_number, address) VALUES (?,?,?,?)`,
		userID,
		strings.TrimSpace(fullname),
		strings.TrimSpace(phonenumber),
		strings.TrimSpace(address),
	)

	if err != nil {
		log.Fatal("Gagal Add Customers",err)
		return
	}

	log.Println("Customers berhasil ditambahkan")
}