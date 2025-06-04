package customer

import (
	"e-commerce-games/entity"
	"e-commerce-games/handler"
	"fmt"
	"strings"
)

func MainMenuCustomer(customer *handler.CustomerHandler, user entity.Customer) {
	var numbersMenu int

	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("Selemat Data %s Sunday Bed Ranger Store\n", user.Name)
	fmt.Println(strings.Repeat("-", 30))

	fmt.Println("1. Order Product")
	fmt.Println("2. List Order")
	fmt.Println("3. Profile")
	fmt.Println("4. Logout")
	fmt.Print("Choose Menu : ")
	fmt.Scan(&numbersMenu)
}