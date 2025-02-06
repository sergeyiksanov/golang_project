package auth_controller

import (
	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/auth_usecase"
	"go.uber.org/zap"
)

type AuthController struct {
	logger  *zap.Logger
	useCase *auth_usecase.AuthUseCase
}

func NewAuthController(logger *zap.Logger, usecase *auth_usecase.AuthUseCase) *AuthController {
	return &AuthController{
		logger:  logger,
		useCase: usecase,
	}
}
