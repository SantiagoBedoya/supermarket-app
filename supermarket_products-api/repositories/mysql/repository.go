package mysql

import (
	"errors"
	"log"

	"github.com/SantiagoBedoya/supermarket_products-api/products"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func newMySQLClient(mysqlURI string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(mysqlURI), &gorm.Config{})
}

func NewMySQLRepository(mysqlURI string) products.Repository {
	db, err := newMySQLClient(mysqlURI)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	db.AutoMigrate(&products.Product{})
	return &repository{db}
}

func (r *repository) Create(product *products.Product) (*products.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		var mysqlErr *mysqlDriver.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, products.ProductAlreadyExistsErr
		}
		return nil, err
	}
	return product, nil
}
func (r *repository) GetAll() ([]products.Product, error) {
	products := make([]products.Product, 0)
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (r *repository) GetByCode(code string) (*products.Product, error) {
	var product products.Product
	if err := r.db.First(&product, "code = ?", code).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, products.ProductNotFoundErr
		}
		return nil, err
	}
	return &product, nil
}
func (r *repository) UpdateByCode(code string, product *products.Product) (*products.Product, error) {
	if err := r.db.Model(product).Where("code = ?", code).Updates(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (r *repository) DeleteByCode(code string) error {
	if err := r.db.Delete(&products.Product{}, "code = ?", code).Error; err != nil {
		return err
	}
	return nil
}
