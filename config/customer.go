package config

import (
	"database/sql"
	"e-commerce-games/entity"
)

func IsEmailUnique(db *sql.DB, email string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func GetCustomerWithUser(db *sql.DB, email string) (*entity.Customer, error)  {
	query := `SELECT c.customer_id, c.full_name, c.phone_number, c.address, c.user_id,
	u.user_id, u.email, u.password, u.role, u.token
	FROM customers c
	JOIN users u ON c.user_id = u.user_id
	WHERE u.email = ?;`

	row := db.QueryRow(query, email)

	var customer entity.Customer
	var user entity.User

	err := row.Scan(
		&customer.ID, &customer.Name, &customer.Phone, &customer.Address, &customer.UserID,
		&user.ID, &user.Email, &user.Password, &user.Role, &user.Token,
	)

	if err != nil {
		return nil, err
	}

	customer.User = &user
	return &customer, nil
}

func UpdateUserToken(db *sql.DB, email string, token string) error {
	query := `UPDATE users SET token = ? WHERE email = ?`
	_, err := db.Exec(query, token, email)
	return err
}