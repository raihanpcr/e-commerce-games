package utils

import "e-commerce-games/entity"

func AddToCart(cart *entity.OrderProduct, product entity.Product, qty int, userID int) {
	itemExists := false

	for i, item := range cart.OrderItem {
		if item.ProductID == product.ProductID {
			// Produk sudah ada, update quantity dan subtotal
			cart.TotalAmount -= cart.OrderItem[i].Subtotal
			cart.OrderItem[i].Quantity += qty
			cart.OrderItem[i].Subtotal = product.Price * cart.OrderItem[i].Quantity
			cart.TotalAmount += cart.OrderItem[i].Subtotal
			itemExists = true
			break
		}
	}

	if !itemExists {
		// Produk belum ada, tambahkan baru
		subtotal := product.Price * qty
		cart.OrderItem = append(cart.OrderItem, entity.OrderItem{
			ProductID: product.ProductID,
			Quantity:  qty,
			UnitPrice: product.Price,
			Subtotal:  subtotal,
		})
		cart.CustomerID = userID
		cart.TotalAmount += subtotal
	}
}
