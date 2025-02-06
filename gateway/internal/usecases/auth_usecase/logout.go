package auth_usecase

import (
	"context"

	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

func (a *AuthUseCase) Logout(ctx context.Context, tokens *models.Tokens) error {
	return a.authAdapter.Logout(ctx, tokens)
}
