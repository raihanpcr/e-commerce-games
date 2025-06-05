package admin

import (
	"bufio"
	"e-commerce-games/entity"
	"e-commerce-games/handler"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func MenuProduct(customer *handler.CustomerHandler, user *entity.Customer) {

	var numbersMenu int
	productHandler := handler.ProductHandler{DB: customer.DB}

	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("Menu Product di Sunday Bed Ranger Store\n")
	for {
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("1. List Data")
		fmt.Println("2. Add Data")
		fmt.Println("0. Kembali")
		fmt.Print("Choose Menu : ")
		fmt.Scan(&numbersMenu)

		switch numbersMenu {
		case 1:
			products, err := productHandler.ListProduct()
			if err != nil {
				log.Fatal("Gagal mengambil data order ", err)
				break
			}
			fmt.Println("Data Order Customers")
			fmt.Println(strings.Repeat("=", 82))
			fmt.Printf("| %-10s | %-19s | %-20s | %-15s | %-8s |\n",
				"Product ID", "Product Name", "Description", "Price", "Stock")
			fmt.Println(strings.Repeat("=", 82))

			for _, p := range products {
				fmt.Printf("| %-10d | %-19s | %-20s | %-15d | %-8d |\n",
					p.ProductID, p.Name, p.Description, p.Price, p.Stock)
			}

			fmt.Println(strings.Repeat("=", 82))
		case 2:
			reader := bufio.NewReader(os.Stdin)
			reader.ReadString('\n')

			fmt.Print("Enter product name: ")
			name, _ := reader.ReadString('\n')

			fmt.Print("Enter product description: ")
			description, _ := reader.ReadString('\n')

			fmt.Print("Enter product price: ")
			priceStr, _ := reader.ReadString('\n')

			fmt.Print("Enter product stock: ")
			stockStr, _ := reader.ReadString('\n')

			name = strings.TrimSpace(name)
			description = strings.TrimSpace(description)
			priceStr = strings.TrimSpace(priceStr)
			stockStr = strings.TrimSpace(stockStr)

			price, errPrice := strconv.Atoi(priceStr)
			stock, errStock := strconv.Atoi(stockStr)

			if errPrice != nil || errStock != nil {
				fmt.Println("Input harga atau stok tidak valid. Harus berupa angka.")
				break
			}

			reqProduct := entity.Product{
				Name:        name,
				Description: description,
				Price:       price,
				Stock:       stock,
			}
			productHandler.AddProduct(reqProduct)

		case 0:
			return
			// MainMenuAdmin(customer, user)
		default:
			fmt.Println("Invalid Input!")
		}
	}
}
