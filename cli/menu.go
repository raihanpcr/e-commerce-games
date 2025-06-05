package cli

import (
	"e-commerce-games/handler"
	"fmt"
	"strings"
)

func MainMenu(userHandler *handler.CustomerHandler) {
	for {
		var numbersMenu int

		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("Selemat Datang di Sunday Bed Ranger Store")
		fmt.Println(strings.Repeat("-", 30))

		//Menu
		fmt.Println("1. List Product")
		fmt.Println("2. Login")
		fmt.Println("3. Register")
		fmt.Print("Choose Menu : ")
		fmt.Scan(&numbersMenu)

		switch numbersMenu {
		case 1:
			productHandler := handler.ProductHandler{DB: userHandler.DB}

			products, err := productHandler.ListProduct()
			if err != nil {
				fmt.Println("Gagal mengambil data produk:", err)
				break
			}
			fmt.Println("List of Products:")
			fmt.Println(strings.Repeat("=", 97))
			fmt.Printf("| %-10s | %-25s | %-30s | %-10s | %-6s |\n",
				"Product ID", "Product Name", "Description", "Price", "Stock")
			fmt.Println(strings.Repeat("=", 97))

			for _, p := range products {
				fmt.Printf("| %-10d | %-25s | %-30s | %-10d | %-6d |\n",
					p.ProductID, p.Name, p.Description, p.Price, p.Stock)
			}

			fmt.Println(strings.Repeat("=", 97))

		case 2:
			//TODO : Login
			LoginUser(userHandler)
		case 3:
			//TODO : Register Customer
			AddUserMenu(userHandler)

		default:
		}
	}
}
