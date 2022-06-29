package accounts

type Service interface {
	SignUp(account *AccountSignUp) (*Account, error)
	SignIn(account *AccountSignIn) (string, error)
	VerifyToken(cookie string) (*Account, error)
}
