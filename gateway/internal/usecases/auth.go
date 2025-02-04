package usecases

import (
	"github.com/sergeyiksanov/golang_project/internal/clients"
	"github.com/sergeyiksanov/golang_project/internal/usecases/models"

	"go.uber.org/zap"
)

type AuthUseCase struct {
	logger     *zap.Logger
	authClient clients.IAuthClient
}

// Logout implements IAuthUseCase.
func (a *AuthUseCase) Logout() {
	panic("unimplemented")
}

// RefreshTokens implements IAuthUseCase.
func (a *AuthUseCase) RefreshTokens() {
	panic("unimplemented")
}

// SignIn implements IAuthUseCase.
func (a *AuthUseCase) SignIn() {
	panic("unimplemented")
}

// SignUp implements IAuthUseCase.
func (a *AuthUseCase) SignUp(credentials *models.Credentials) (*models.Tokens, error) {
	tokens, err := a.authClient.SignUp(credentials)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func NewAuthUseCase(logger *zap.Logger, authClient clients.IAuthClient) IAuthUseCase {
	return &AuthUseCase{
		logger:     logger,
		authClient: authClient,
	}
}
