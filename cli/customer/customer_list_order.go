package customer

import (
	"database/sql"
	"e-commerce-games/entity"
	"e-commerce-games/handler"
	"fmt"
	"log"
	"strings"
)

func ListOrderCustomer(customer *handler.CustomerHandler, user *entity.Customer, orderProduct *handler.OrderProductHandler, db *sql.DB) {
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("List Order Customer ", user.Name, " di Sunday Bed Ranger Store")
	fmt.Println(strings.Repeat("-", 30))

	// Ambil data order berdasarkan customer
	dataOrderProduct, err := orderProduct.ListOrderByCustomer(*user)
	if err != nil {
		log.Println("Gagal mengambil data order:", err)
		fmt.Println("Terjadi kesalahan saat mengambil data order.")
		return
	}

	if len(dataOrderProduct) == 0 {
		fmt.Println("Belum ada order yang dibuat.")
		return
	}

	for _, order := range dataOrderProduct {
		fmt.Println(strings.Repeat("-", 50))
		fmt.Printf("Order ID     : %d\n", order.OrderID)
		fmt.Printf("Tanggal      : %s\n", order.OrderDate)
		fmt.Printf("Status       : %s\n", order.Status)
		fmt.Printf("Total Amount : Rp%d\n", order.TotalAmount)
		fmt.Println("Items:")
		fmt.Printf("%-4s %-20s %-10s %-10s %-10s\n", "No", "Nama Produk", "Harga", "Qty", "Subtotal")

		for i, item := range order.OrderItem {
			fmt.Printf("%-4d %-20s Rp%-9d %-10d Rp%-10d\n",
				i+1,
				item.Product.Name,
				item.UnitPrice,
				item.Quantity,
				item.Subtotal,
			)
		}
		fmt.Println(strings.Repeat("-", 50))
	}
}
