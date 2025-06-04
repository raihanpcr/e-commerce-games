package customer

import (
	"database/sql"
	"e-commerce-games/entity"
	"e-commerce-games/handler"
	"fmt"
	"log"
	"strings"
)

func OrderProduct(customer *handler.CustomerHandler, user entity.Customer, products *handler.ProductHandler, db *sql.DB) {
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Order Product di E-Commerce Games")
	fmt.Println(strings.Repeat("-", 30))

	//Menu
	dataProduct, err := products.ListProduct()
	if err != nil {
		log.Fatal("Failed get products:", err)
		fmt.Println("Failed get products")
		return
	}
	for _, product := range dataProduct {
		fmt.Printf("%d. %s (%d)\n", product.ProductID, product.Name, product.Price)
	}

	cart := entity.OrderProduct{
		OrderItem:   []entity.OrderItem{},
		TotalAmount: 0,
	}

	for {
		var productID int
		var productQuantity int

		fmt.Print("Choose id product : ")
		fmt.Scan(&productID)

		found := false
		for _, product := range dataProduct {
			if productID == product.ProductID {
				for {
					fmt.Printf("How many units?")
					fmt.Scan(&productQuantity)

					//error handling 3
					if productQuantity > 0 && productQuantity <= 100 { //true
						break
					} else { //false
						fmt.Printf("Number of unit too large")
					}
				}

				// addToChart
				subtotal := product.Price * productQuantity
				cart.OrderItem = append(cart.OrderItem, entity.OrderItem{
					ProductID: product.ProductID,
					Quantity:  productQuantity,
					UnitPrice: product.Price,
					Subtotal:  subtotal,
				})
				cart.CustomerID = user.ID
				cart.TotalAmount += subtotal

				fmt.Println("Order Summary:")
				for _, item := range cart.OrderItem {
					fmt.Printf("- Product ID: %d | Qty: %d | Subtotal: Rp%d\n", item.ProductID, item.Quantity, item.Subtotal)
				}
				fmt.Printf("Total: Rp%d\n", cart.TotalAmount)

				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Item not available. Please select a valid product form the menu\n")
		}

		var confirm string
		fmt.Printf("Would you like to order another product? (yes/no)")
		fmt.Scan(&confirm)
		if confirm == "no" {
			h := handler.OrderProductHandler{DB: db}
			// order, err :=
			h.AddOderProduct(cart) //get order id
			// if err != nil {
			// 	return
			// }
			break
		}

	}

	var paymentMethod int
	var confirm string
	for {
		fmt.Println("Choose payment method:")
		fmt.Println("1. Cash On Delivery (COD)")
		fmt.Println("2. Bank Transfer")
		fmt.Print("Enter method (1/2): ")
		fmt.Scan(&paymentMethod)

		if paymentMethod == 1 {
			// add payment with order id
			fmt.Println("Payment will be made at delivery. Thank you for your order!")
			break
		} else if paymentMethod == 2 {
			for {
				fmt.Println("Please transfer to bank number 0202020")
				fmt.Print("Confirm payment? (yes/no): ")
				fmt.Scan(&confirm)
				if strings.ToLower(confirm) == "yes" {
					// add payment with order id
					fmt.Println("Thank you for your order and payment!")
					break
				} else {
					fmt.Println("Please complete the payment.")
				}
			}
			break
		} else {
			fmt.Println("Invalid option. Please choose 1 or 2.")
		}
	}

	for {
		fmt.Printf("enter 'back' for back to menu: ")
		fmt.Scan(&confirm)
		if strings.ToLower(confirm) == "back" {
			MainMenuCustomer(customer, user)
			break
		} else {
			fmt.Println("Invalid option!")
		}
	}
}
