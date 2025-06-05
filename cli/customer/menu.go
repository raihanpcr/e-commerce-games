package customer

import (
	"e-commerce-games/config"
	"e-commerce-games/entity"
	"e-commerce-games/handler"
	"fmt"
	"strings"
)

func MainMenuCustomer(customer *handler.CustomerHandler, user *entity.Customer) {
	db := config.InitDB()
	var numbersMenu int
	// var isLoop = true

	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("Selemat Data %s Sunday Bed Ranger Store\n", user.Name)
	for {
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("1. Order Product")
		fmt.Println("2. List Order")
		fmt.Println("3. Profile")
		fmt.Println("4. Logout")
		fmt.Print("Choose Menu : ")
		fmt.Scan(&numbersMenu)

		switch numbersMenu {
		case 1:
			OrderProduct(customer, user, &handler.ProductHandler{DB: db}, db)
		case 2:
			ListOrderCustomer(customer, user, &handler.OrderProductHandler{DB: db}, db)
		case 3:
			fmt.Println(strings.Repeat("-", 30))
			fmt.Println("Profile Anda")
			fmt.Println(strings.Repeat("-", 30))

			fmt.Println("Customer Name :", user.Name)
			fmt.Println("Customer Address :", user.Address)
			fmt.Println("Customer Phone :", user.Phone)

			//TODO : Update data customer
		case 4:
			//todo : clear token
			err := config.ClearUserToken(customer.DB, user.User.Email)
			if err != nil {
				fmt.Println("Gagal logout:", err)
			}
			fmt.Println("Logout berhasil.")

			//todo : clear entity customer
			*user = entity.Customer{}
			return
		}
	}

}
