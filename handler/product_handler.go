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

func (h *ProductHandler) ListBestSellingProduct() ([]entity.BestSellingProduct, error) {
	query := `
		SELECT 
			p.product_id,
			p.name AS product_name,
			SUM(oi.quantity) AS Terjual
		FROM orderitems oi
		JOIN products p ON oi.product_id = p.product_id
		GROUP BY p.product_id, p.name
		ORDER BY Terjual DESC
		LIMIT 3
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil produk terlaris: %w", err)
	}
	defer rows.Close()

	var bestSellers []entity.BestSellingProduct

	for rows.Next() {
		var item entity.BestSellingProduct
		if err := rows.Scan(&item.ProductID, &item.ProductName, &item.Terjual); err != nil {
			log.Printf("gagal scan produk terlaris: %v\n", err)
			continue
		}
		bestSellers = append(bestSellers, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("kesalahan saat membaca hasil: %w", err)
	}

	return bestSellers, nil
}
