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
	order, err := h.DB.Exec(queryOrder, reqOrderProduct.CustomerID, now, "unpaid", reqOrderProduct.TotalAmount)
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

func (h *OrderProductHandler) ListOrderByCustomer(customer entity.Customer) ([]entity.OrderProduct, error) {
	query := `
		SELECT
			o.order_id,
			o.customer_id,
			o.order_date,
			o.total_amount,
			o.status,
			oi.order_item_id,
			oi.product_id,
			p.name as product_name,
			p.description,
			p.price,
			p.stock,
			oi.unit_price,
			oi.quantity,
			oi.total_amount as subtotal
		FROM
			orders o
		JOIN customers c ON o.customer_id = c.customer_id
		JOIN orderitems oi ON oi.order_id = o.order_id
		JOIN products p ON oi.product_id = p.product_id
		WHERE
			o.customer_id = ?
		ORDER BY 
			o.order_id DESC, oi.order_item_id ASC`

	rows, err := h.DB.Query(query, customer.ID)
	if err != nil {
		log.Println("Query failed:", err)
		return nil, err
	}
	defer rows.Close()

	orderMap := make(map[int]*entity.OrderProduct)
	for rows.Next() {
		var (
			orderID, customerID, totalAmount, orderItemID, productID, unitPrice, quantity, subtotal, price, stock int
			orderDate, status, productName, productDesc                                                           string
		)

		err := rows.Scan(
			&orderID, &customerID, &orderDate, &totalAmount, &status,
			&orderItemID, &productID, &productName, &productDesc, &price, &stock,
			&unitPrice, &quantity, &subtotal,
		)
		if err != nil {
			log.Println("Scan failed:", err)
			continue
		}

		item := entity.OrderItem{
			OrderItemID: orderItemID,
			OrderID:     orderID,
			ProductID:   productID,
			Quantity:    quantity,
			UnitPrice:   unitPrice,
			Subtotal:    subtotal,
			Product: entity.Product{
				ProductID:   productID,
				Name:        productName,
				Description: productDesc,
				Price:       price,
				Stock:       stock,
			},
		}

		// Jika order belum ada, buat entry baru
		if _, exists := orderMap[orderID]; !exists {
			orderMap[orderID] = &entity.OrderProduct{
				OrderID:     orderID,
				CustomerID:  customerID,
				OrderDate:   orderDate,
				TotalAmount: totalAmount,
				Status:      status,
				OrderItem:   []entity.OrderItem{},
			}
		}

		// Tambahkan item ke order yang sesuai
		orderMap[orderID].OrderItem = append(orderMap[orderID].OrderItem, item)
	}

	var orders []entity.OrderProduct
	for _, v := range orderMap {
		orders = append(orders, *v)
	}

	return orders, nil
}
