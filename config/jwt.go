package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func init() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Gagal load file .env file : ", err)
	}

	jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateJWT(email, role string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ClearUserToken(db *sql.DB, email string) error {
	_, err := db.Exec(`UPDATE users SET token = NULL WHERE email = ?`, email)
	return err
}