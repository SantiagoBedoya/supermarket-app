package main

import (
	"log"
	"os"

	"github.com/SantiagoBedoya/supermarket_accounts-api/accounts"
	"github.com/SantiagoBedoya/supermarket_accounts-api/repositories/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	repository := mysql.NewMySQLRepository(os.Getenv("MYSQL_URI"))

	accounts := []accounts.Account{
		{
			FirstName: "Santiago",
			LastName:  "Bedoya",
			Email:     "santiago@gmail.com",
			Password:  encryptPassword("test1234"),
			Role:      "admin",
		},
	}
	for _, account := range accounts {
		_, err := repository.Create(&account)
		if err != nil {
			log.Printf("error creating user %s: %v\n", account.Email, err)
		}
	}
	log.Println("account seeder succeeded")
}

func encryptPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
