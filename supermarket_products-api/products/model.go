package products

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Category int64

const (
	Food     Category = 1
	Drink             = 2
	Personal          = 3
	Clean             = 4
	EService          = 5
)

type Product struct {
	gorm.Model
	Name     string   `json:"name" gorm:"notNull" validate:"required"`
	Code     string   `json:"code" gorm:"notNull;unique" validate:"required"`
	Stock    uint     `json:"stock" gorm:"notNull;default:0"`
	Category Category `json:"category" gorm:"notNull" validate:"required"`
	Price    float64  `json:"price" gorm:"notNull" validate:"required"`
}

func (p *Product) ToPublicProduct() *PublicProduct {
	return &PublicProduct{
		Name:     p.Name,
		Code:     p.Code,
		Stock:    p.Stock,
		Category: p.Category,
		Price:    p.Price,
	}
}

type PublicProduct struct {
	Name     string   `json:"name" gorm:"notNull" validate:"required"`
	Code     string   `json:"code" gorm:"notNull;unique" validate:"required"`
	Stock    uint     `json:"stock" gorm:"notNull;default:0"`
	Category Category `json:"category" gorm:"notNull" validate:"required"`
	Price    float64  `json:"price" gorm:"notNull" validate:"required"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return InvalidProductDataErr
	}
	return nil
}
