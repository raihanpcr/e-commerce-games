package main

import (
	"e-commerce-games/cli"
	"e-commerce-games/config"
	"e-commerce-games/handler"
	"fmt"
	"strings"
)

func main() {

	db := config.InitDB()

	var numbersMenu int

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Selemat Data di E-Commerce Games")
	fmt.Println(strings.Repeat("-", 30))

	//Menu
	fmt.Println("1. List Product")
	fmt.Println("2. Signin")
	fmt.Println("3. Signup")
	fmt.Print("Choose Menu : ")
	fmt.Scan(&numbersMenu)

	switch numbersMenu {
	case 1:
	case 2:
		//sign up
		cli.AddUserMenu(&handler.CustomerHandler{DB: db})
	case 3:

	default:
	}
}