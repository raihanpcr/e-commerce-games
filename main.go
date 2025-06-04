package main

import (
	"e-commerce-games/cli"
	"e-commerce-games/config"
	"e-commerce-games/handler"
)

func main() {

	db := config.InitDB()

	cli.MainMenu(&handler.CustomerHandler{DB: db})

	// cli.CustomerMenu(db)
}
