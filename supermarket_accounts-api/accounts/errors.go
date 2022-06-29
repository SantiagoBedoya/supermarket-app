package accounts

import (
	"errors"
)

var (
	InvalidUserSignUpErr   = errors.New("invalid user signUp data")
	InvalidUserSignInErr   = errors.New("invalid user signIn data")
	InvalidUserCredentials = errors.New("invalid user credentials")
	UserAlreadyExistErr    = errors.New("this email is already in use")
	UserUnauthorizedErr    = errors.New("Unauthorized user")
)
