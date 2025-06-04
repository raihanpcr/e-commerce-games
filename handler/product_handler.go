package handler

import (
	"database/sql"
	"e-commerce-games/entity"
	"fmt"
	"log"
)

type ProductHandler struct {
	DB *sql.DB
}

func (h *ProductHandler) AddProduct(reqProduct entity.Product) {
	query := fmt.Sprint("INSERT INTO products (name,description,price,stock) VALUES (?,?,?,?)")
	_, err := h.DB.Exec(query, reqProduct.Name, reqProduct.Description, reqProduct.Price, reqProduct.Stock)
	if err != nil {
		log.Fatal("Store Product failed, ", err)
		fmt.Println("Store Product failed")
		return
	}

	fmt.Println("Store Product Product successfully")
}

func (h *ProductHandler) ListProduct() ([]entity.Product, error) {
	query := fmt.Sprint("SELECT product_id, name, description, price, stock FROM products")
	rows, err := h.DB.Query(query)
	if err != nil {
		log.Fatal("View Product failed, ", err)
		fmt.Println("View Product failed")
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var item entity.Product
		err := rows.Scan(&item.ProductID, &item.Name, &item.Description, &item.Price, &item.Stock)
		if err != nil {
			log.Println("Failed to scan product:", err)
			continue
		}
		products = append(products, item)
	}
	return products, nil
}
