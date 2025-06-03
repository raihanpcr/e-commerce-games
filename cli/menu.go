package cli

import (
	"e-commerce-games/handler"
	"fmt"
	"strings"
)

func MainMenu(userHandler *handler.CustomerHandler) {

	var numbersMenu int

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Selemat Data di E-Commerce Games")
	fmt.Println(strings.Repeat("-", 30))

	//Menu
	fmt.Println("1. List Product")
	fmt.Println("2. Signup")
	fmt.Println("3. Signin")
	fmt.Print("Choose Menu : ")
	fmt.Scan(&numbersMenu)

	switch numbersMenu {
	case 1:
	case 2:
		//TODO : Sign up customer
		AddUserMenu(userHandler)
	case 3:

	default:
	}
}