package auth_usecase

import (
	"context"

	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
	"go.uber.org/zap"
)

type (
	authAdapter interface {
		SignUp(ctx context.Context, credentials *models.Credentials) error
		SignIn(ctx context.Context, credentials *models.Credentials) (*models.Tokens, error)
		RefreshTokens(ctx context.Context, refreshToken models.RefreshToken) (*models.Tokens, error)
		Logout(ctx context.Context, tokens *models.Tokens) error
	}

	AuthUseCase struct {
		logger      *zap.Logger
		authAdapter authAdapter
	}
)

func NewAuthUseCase(logger *zap.Logger, authAdapter authAdapter) *AuthUseCase {
	return &AuthUseCase{
		logger:      logger,
		authAdapter: authAdapter,
	}
}
