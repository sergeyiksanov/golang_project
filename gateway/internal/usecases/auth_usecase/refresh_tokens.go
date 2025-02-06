package auth_usecase

import (
	"context"

	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

func (a *AuthUseCase) RefreshTokens(ctx context.Context, refreshToken models.RefreshToken) (*models.Tokens, error) {
	tokens, err := a.authAdapter.RefreshTokens(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
