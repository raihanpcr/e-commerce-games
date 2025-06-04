package customer

import (
	"e-commerce-games/config"
	"e-commerce-games/entity"
	"e-commerce-games/handler"
	"fmt"
	"strings"
)

func MainMenuCustomer(customer *handler.CustomerHandler, user entity.Customer) {
	db := config.InitDB()
	var numbersMenu int

	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("Selemat Data %s Sunday Bed Ranger Store\n", user.Name)
	fmt.Println(strings.Repeat("-", 30))

	//Menu
	fmt.Println("1. Order Product")
	fmt.Println("2. list Order")
	fmt.Println("3. Profile")
	fmt.Println("4. Logout")
	fmt.Print("Choose Menu : ")
	fmt.Scan(&numbersMenu)

	switch numbersMenu {
	case 1:
		OrderProduct(customer, user, &handler.ProductHandler{DB: db}, db)
	case 2:
	case 3:
	case 4:
	default:
	}
}
