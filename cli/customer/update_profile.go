package customer

import (
	"bufio"
	"e-commerce-games/entity"
	"e-commerce-games/handler"
	"fmt"
	"os"
	"strings"
)

func UpdateProfileCustomer(customer *handler.CustomerHandler, cust *entity.Customer) {
	
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Update Profile Anda")
	fmt.Println(strings.Repeat("-", 30))

	reader := bufio.NewReader(os.Stdin)
	// reader.ReadString('\n')

	//input full_name
	fmt.Print("Update your Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Update your Phone Number: ")
	phone, _ := reader.ReadString('\n')

	fmt.Print("Update your Address: ")
	address, _ := reader.ReadString('\n')

	cust.Name= strings.TrimSpace(name)
	cust.Phone = strings.TrimSpace(phone)
	cust.Address = strings.TrimSpace(address)

	customer.UpdateCustomer(cust)
}