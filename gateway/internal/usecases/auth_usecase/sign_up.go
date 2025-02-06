package auth_usecase

import (
	"context"

	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

func (a *AuthUseCase) SignUp(ctx context.Context, credentials *models.Credentials) error {
	err := a.authAdapter.SignUp(ctx, credentials)
	if err != nil {
		return err
	}

	return nil
}
