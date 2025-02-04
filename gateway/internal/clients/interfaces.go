package clients

import "github.com/sergeyiksanov/golang_project/internal/usecases/models"

type IAuthClient interface {
	SignUp(credentials *models.Credentials) (*models.Tokens, error)
	SignIn()
	VerifyAccessToken()
	RefreshTokens()
	Logout()
}
