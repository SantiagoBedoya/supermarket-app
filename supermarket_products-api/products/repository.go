package products

type Repository interface {
	Create(product *Product) (*Product, error)
	GetAll() ([]Product, error)
	GetByCode(code string) (*Product, error)
	UpdateByCode(code string, product *Product) (*Product, error)
	DeleteByCode(code string) error
}
