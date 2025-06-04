package handler

import (
	"database/sql"
	"e-commerce-games/entity"
	"fmt"
	"log"
	"time"
)

type PaymentHandler struct {
	DB *sql.DB
}

func (h *PaymentHandler) AddPayment(reqPayment entity.Payment) {
	now := time.Now().Format("2006-01-02")

	queryPayment := fmt.Sprint("INSERT INTO payments (order_id,payment_method,payment_date,payment_paid) VALUES (?,?,?,?)")
	_, err := h.DB.Exec(queryPayment, reqPayment.OrderID, reqPayment.PaymentMethod, now, reqPayment.PaymentPaid)
	if err != nil {
		log.Fatal("Store Payment failed, ", err)
		return
	}

	// update stok produk
	_, err = h.DB.Exec("UPDATE orders SET status = ? WHERE order_id = ?", "paid", reqPayment.OrderID)
	if err != nil {
		log.Fatal("Failed to update order status: ", err)
		return
	}

	fmt.Println("Store Payment successfully")
}
