package admin

import (
	"e-commerce-games/config"
	"e-commerce-games/entity"
	"e-commerce-games/handler"
	"fmt"
	"log"
	"strings"
)

func MainMenuAdmin(customer *handler.CustomerHandler, user *entity.Customer) {
	
	var numbersMenu int
	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("Selemat Data %s Sunday Bed Ranger Store\n", user.Name)
	for {
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("1. Product")
		fmt.Println("2. Lihat Order")
		fmt.Println("3. Report")
		fmt.Println("4. Logout")
		fmt.Print("Choose Menu : ")
		fmt.Scan(&numbersMenu)

		switch numbersMenu {
		case 1:
		case 2:
			orderHandler := handler.OrderHandler{DB: customer.DB}

			orders, err := orderHandler.ListOrder()
			if err != nil {
				log.Fatal("Gagal mengambil data order ",err)
				break
			}
			fmt.Println("Data Order Customers")
			fmt.Println(strings.Repeat("=", 100))
			fmt.Printf("| %-10s | %-19s | %-20s | %-15s | %-8s | %-10s |\n",
				"Order ID", "Tanggal Order", "Customer", "Product", "Qty", "Amount")
			fmt.Println(strings.Repeat("=", 100))

			for _, o := range orders {
				fmt.Printf("| %-10d | %-19s | %-20s | %-15s | %-8d | %-10d |\n",
					o.OrderID, o.OrderDate, o.FullName, o.ProductName, o.Quantity, o.TotalAmount)
			}

			fmt.Println(strings.Repeat("=", 100))

		case 3:
			fmt.Println(strings.Repeat("-", 30))
			productHandler := handler.ProductHandler{DB: customer.DB}

			customerHandler := handler.CustomerHandler{DB: customer.DB}
			var reportNumbers int
			fmt.Println("1. Stok Product Sedikit")
			fmt.Println("2. Pendapatan Perbulan")
			fmt.Println("3. Product yang paling laku")
			//TODO 
			fmt.Println("4. Customer terbanyak Order")
			fmt.Print("Choose Report : ")
			fmt.Scan(&reportNumbers)

			switch reportNumbers {
			case 1:
			case 2:
			case 3:
				products, err := productHandler.ListBestSellingProduct()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(strings.Repeat("=", 60))
				fmt.Printf("| %-10s | %-30s | %-10s |\n", "Product ID", "Product Name", "Terjual")
				fmt.Println(strings.Repeat("=", 60))
				for _, p := range products {
					fmt.Printf("| %-10d | %-30s | %-10d |\n", p.ProductID, p.ProductName, p.Terjual)
				}
				fmt.Println(strings.Repeat("=", 60))
			case 4:
				customers, err := customerHandler.ListTopCustomers()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(strings.Repeat("=", 65))
				fmt.Printf("| %-12s | %-30s | %-12s |\n", "Customer ID", "Customer Name", "Total Orders")
				fmt.Println(strings.Repeat("=", 65))
				for _, c := range customers {
					fmt.Printf("| %-12d | %-30s | %-12d |\n", c.CustomerID, c.CustomerName, c.TotalOrders)
				}
				fmt.Println(strings.Repeat("=", 65))
			}
		case 4:
			err := config.ClearUserToken(customer.DB, user.User.Email)
			if err != nil {
				fmt.Println("Gagal logout:", err)
			}
			fmt.Println("Logout berhasil.")

			*user = entity.Customer{}
			return
		}
	}
}