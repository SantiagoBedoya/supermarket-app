package accounts

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) VerifyToken(cookie string) (*Account, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, UserUnauthorizedErr
	}
	claims := token.Claims.(*jwt.StandardClaims)

	account, err := s.repository.FindById(claims.Id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *service) SignUp(account *AccountSignUp) (*Account, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	acc := Account{
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
		Password:  string(hash),
	}

	return s.repository.Create(&acc)
}

func (s *service) SignIn(account *AccountSignIn) (string, error) {
	acc := Account{
		Email: account.Email,
	}
	currentAccount, err := s.repository.FindByEmail(&acc)
	if err != nil {
		return "", InvalidUserCredentials
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(currentAccount.Password),
		[]byte(account.Password),
	); err != nil {
		return "", InvalidUserCredentials
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", currentAccount.ID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return token, nil
}
