package products

type Service interface {
	Create(product *Product) (*Product, error)
	GetAll() ([]PublicProduct, error)
	GetByCode(code string) (*PublicProduct, error)
	UpdateByCode(code string, product *Product) (*Product, error)
	DeleteByCode(code string) error
}
