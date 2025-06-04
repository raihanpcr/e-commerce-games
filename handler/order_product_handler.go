package handler

import (
	"database/sql"
	"e-commerce-games/entity"
	"fmt"
	"log"
	"time"
)

type OrderProductHandler struct {
	DB *sql.DB
}

func (h *OrderProductHandler) AddOderProduct(reqOrderProduct entity.OrderProduct) (int64, error) {
	now := time.Now().Format("2006-01-02")

	queryOrder := fmt.Sprint("INSERT INTO orders (customer_id,order_date,status,total_amount) VALUES (?,?,?,?)")
	order, err := h.DB.Exec(queryOrder, 1, now, "unpaid", reqOrderProduct.TotalAmount)
	if err != nil {
		log.Fatal("Store Order failed, ", err)
		return 0, err
	}

	orderID, err := order.LastInsertId()
	if err != nil {
		log.Fatal("Failed get order_id:", err)
		fmt.Println("Internal Server Error")
		return 0, err
	}

	for _, v := range reqOrderProduct.OrderItem {
		// cek stok produk
		var currentStock int
		err := h.DB.QueryRow("SELECT stock FROM products WHERE product_id = ?", v.ProductID).Scan(&currentStock)
		if err != nil {
			log.Fatal("Failed to check product stock: ", err)
			return 0, err
		}

		// validasi stok cukup
		if v.Quantity > currentStock {
			log.Fatalf("Stok produk %d tidak cukup (tersedia: %d, diminta: %d)", v.ProductID, currentStock, v.Quantity)
			return 0, err
		}

		// insert
		queryOrderItem := fmt.Sprint("INSERT INTO orderitems (order_id,product_id,quantity,unit_price,total_amount) VALUES (?,?,?,?,?)")
		_, err = h.DB.Exec(queryOrderItem, orderID, v.ProductID, v.Quantity, v.UnitPrice, v.Subtotal)
		if err != nil {
			log.Fatal("Store Order Items failed, ", err)
			fmt.Println("Store Order Product failed")
			return 0, err
		}

		// update stok produk
		_, err = h.DB.Exec("UPDATE products SET stock = stock - ? WHERE product_id = ?", v.Quantity, v.ProductID)
		if err != nil {
			log.Fatal("Failed to update product stock: ", err)
			return 0, err
		}
	}

	fmt.Println("Store Order Product successfully")
	return orderID, nil
}
