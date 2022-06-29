package accounts

type Repository interface {
	Create(account *Account) (*Account, error)
	FindById(userId string) (*Account, error)
	FindByEmail(account *Account) (*Account, error)
}
