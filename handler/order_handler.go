package handler

import (
	"database/sql"
	"e-commerce-games/entity"
	"fmt"
	"log"
)

type OrderHandler struct {
	DB *sql.DB
}

func (h *OrderHandler) ListOrder() ([]entity.OrderDetail, error) {
	query := `
		SELECT 
			o.order_id,
			o.order_date,
			c.full_name AS user_name,
			p.name AS product_name,
			oi.quantity,
			oi.total_amount
		FROM orders o
		JOIN customers c ON o.customer_id = c.customer_id
		JOIN orderitems oi ON o.order_id = oi.order_id
		JOIN products p ON oi.product_id = p.product_id
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data order: %v", err)
	}
	defer rows.Close()

	var orders []entity.OrderDetail

	for rows.Next() {
		var order entity.OrderDetail
		err := rows.Scan(
			&order.OrderID,
			&order.OrderDate,
			&order.FullName,
			&order.ProductName,
			&order.Quantity,
			&order.TotalAmount,
		)
		if err != nil {
			log.Println("gagal membaca data:", err)
			continue
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
