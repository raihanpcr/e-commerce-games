package handler

import (
	"e-commerce-games/config"
	"e-commerce-games/entity"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)



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

func (h *CustomerHandler) UpdateCustomer(fullname, phone, address string, id int) error {

	query := `
		UPDATE customers 
		SET full_name = ?, phone_number = ?, address = ? 
		WHERE customer_id = ?
	`

	_, err := h.DB.Exec(query, fullname, phone, address, id)
	if err != nil {
		return fmt.Errorf("failed to update customer: %v", err)
	}
	
	fmt.Println("Data Sukses Update")
	return nil
}

func (h *CustomerHandler) Login(email, password string) (*entity.Customer, error){
	
	//Get Customer by email
	user, err := config.GetCustomerWithUser(h.DB, email)
	// fmt.Println(user, err)
	if err != nil {
		return nil, fmt.Errorf("email dan password tidak valid: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.User.Password), []byte(password))
	// fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("email dan password tidak valid: %w", err)
	}
	
	//Generate JWT
	token, err := config.GenerateJWT(email, user.User.Role)
	if err != nil {
		return nil, fmt.Errorf("gagal generate token: %w", err)
	}
	
	//Save token
	err = config.UpdateUserToken(h.DB, email, token)
	if err != nil {
		return nil, fmt.Errorf("gagal menyimpan token ke database: %w", err)
	}
	
	fmt.Println("Login Berhasil")
	return user,nil
}

func (h *CustomerHandler) ListTopCustomers() ([]entity.MostOrderCustomer, error) {
	query := `
		SELECT 
			c.customer_id,
			c.full_name AS customer_name,
			COUNT(o.order_id) AS total_orders
		FROM customers c
		JOIN orders o ON c.customer_id = o.customer_id
		GROUP BY c.customer_id, c.full_name
		ORDER BY total_orders DESC
		LIMIT 3
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data customer: %w", err)
	}
	defer rows.Close()

	var topCustomers []entity.MostOrderCustomer

	for rows.Next() {
		var c entity.MostOrderCustomer
		if err := rows.Scan(&c.CustomerID, &c.CustomerName, &c.TotalOrders); err != nil {
			log.Printf("gagal scan data customer: %v", err)
			continue
		}
		topCustomers = append(topCustomers, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error saat membaca baris: %w", err)
	}

	return topCustomers, nil
}
