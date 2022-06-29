package accounts

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"first_name;notNull"`
	LastName  string `json:"last_name" gorm:"last_name;notNull"`
	Email     string `json:"email" gorm:"email;notNull;unique"`
	Password  string `json:"-" gorm:"password;notNull"`
	Role      string `json:"role" gorm:"default:standard"`
}

func (a *Account) ToPublicAccount() *PublicAccount {
	return &PublicAccount{
		FirstName: a.FirstName,
		LastName:  a.LastName,
		Email:     a.Email,
		Role:      a.Role,
	}
}

type PublicAccount struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type AccountSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountSignUp struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
