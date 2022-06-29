package products

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Create(product *Product) (*Product, error) {
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return s.repository.Create(product)
}
func (s *service) GetAll() ([]PublicProduct, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	publicProducts := make([]PublicProduct, 0)
	for _, product := range products {
		publicProducts = append(publicProducts, *product.ToPublicProduct())
	}
	return publicProducts, nil
}
func (s *service) GetByCode(code string) (*PublicProduct, error) {
	product, err := s.repository.GetByCode(code)
	if err != nil {
		return nil, err
	}
	return product.ToPublicProduct(), nil
}
func (s *service) UpdateByCode(code string, product *Product) (*Product, error) {
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return s.repository.UpdateByCode(code, product)
}
func (s *service) DeleteByCode(code string) error {
	return s.repository.DeleteByCode(code)
}
