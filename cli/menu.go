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
		fmt.Println("Selemat Data di E-Commerce Games")
		fmt.Println(strings.Repeat("-", 30))

		//Menu
		fmt.Println("1. List Product")
		fmt.Println("2. Login")
		fmt.Println("3. Register")
		fmt.Print("Choose Menu : ")
		fmt.Scan(&numbersMenu)

		switch numbersMenu {
		case 1:
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