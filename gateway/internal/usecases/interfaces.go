package usecases

import "github.com/sergeyiksanov/golang_project/internal/usecases/models"

type IAuthUseCase interface {
	SignIn()
	SignUp(credentials *models.Credentials) (*models.Tokens, error)
	RefreshTokens()
	Logout()
}
