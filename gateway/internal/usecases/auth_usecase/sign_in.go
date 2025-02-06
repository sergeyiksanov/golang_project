package auth_usecase

import (
	"context"

	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

func (a *AuthUseCase) SignIn(ctx context.Context, credentials *models.Credentials) (*models.Tokens, error) {
	tokens, err := a.authAdapter.SignIn(ctx, credentials)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
