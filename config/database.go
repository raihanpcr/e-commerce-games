package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() *sql.DB{

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal koneksi ke database : ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database tidak bisa diakses : ", err)
	}

	return db

}