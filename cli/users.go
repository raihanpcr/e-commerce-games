package cli

import (
	"bufio"
	"e-commerce-games/config"
	"e-commerce-games/handler"
	"fmt"
	"log"
	"os"
	"strings"
)

func AddUserMenu(userHandler *handler.CustomerHandler) {

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Selamat Datang di Menu Pendaftaran")
	fmt.Println(strings.Repeat("-", 30))

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	var email string
	for {
		//input email
		fmt.Print("Enter your email: ")
		emailInput, _ := reader.ReadString('\n')
		email = strings.TrimSpace(emailInput)

		//cek email apakah sudah digunakan
		isEmailUnique, err := config.IsEmailUnique(userHandler.DB, email)

		if err != nil {
			log.Fatal("Database error:", err)
		}
		if !isEmailUnique {
			fmt.Println("Email sudah digunakan, coba yang lain.")
			continue
		}
		break
	}
	
	//input password
	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')

	//input full_name
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter your Phone Number: ")
	phone, _ := reader.ReadString('\n')

	fmt.Print("Enter your Address: ")
	address, _ := reader.ReadString('\n')

	
	password = strings.TrimSpace(password)
	name = strings.TrimSpace(name)
	phone = strings.TrimSpace(phone)
	address = strings.TrimSpace(address)

	userHandler.AddCustomer(email, password, name, phone, address)
}