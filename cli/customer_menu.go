package cli

import (
	"database/sql"
	"e-commerce-games/handler"
	"fmt"
	"strings"
)

func CustomerMenu(db *sql.DB) {

	var numbersMenu int

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Selemat Data di E-Commerce Games")
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
		OrderProduct(&handler.ProductHandler{DB: db}, db)
	case 2:
	case 3:
	case 4:
	default:
	}
}
