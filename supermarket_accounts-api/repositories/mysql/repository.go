package mysql

import (
	"errors"
	"log"

	mysqlDriver "github.com/go-sql-driver/mysql"

	"github.com/SantiagoBedoya/supermarket_accounts-api/accounts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func newMySQLClient(mysqlURI string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(mysqlURI), &gorm.Config{})

}

func NewMySQLRepository(mysqlURI string) accounts.Repository {
	db, err := newMySQLClient(mysqlURI)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	db.AutoMigrate(&accounts.Account{})
	return &repository{db}
}

func (r *repository) Create(account *accounts.Account) (*accounts.Account, error) {
	if err := r.db.Create(account).Error; err != nil {
		var mysqlErr *mysqlDriver.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, accounts.UserAlreadyExistErr
		}
		return nil, err
	}
	return account, nil
}

func (r *repository) FindById(userId string) (*accounts.Account, error) {
	account := &accounts.Account{}
	if err := r.db.First(account, "id = ?", userId).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *repository) FindByEmail(account *accounts.Account) (*accounts.Account, error) {
	if err := r.db.First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
