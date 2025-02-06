package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sergeyiksanov/golang_project/gateway/internal/api/controllers/auth_controller"
	"go.uber.org/zap"
)

type AuthRoutes struct {
	logger     *zap.Logger
	handler    *gin.Engine
	controller *auth_controller.AuthController
}

func NewAuthRoutes(logger *zap.Logger, handler *gin.Engine, controller *auth_controller.AuthController) *AuthRoutes {
	return &AuthRoutes{
		logger:     logger,
		handler:    handler,
		controller: controller,
	}
}

func (ar AuthRoutes) Setup() {
	ar.logger.Info(
		"Setup auth routes",
		zap.String("version", "1.0.0"),
		zap.Int("port", 8080),
	)

	auth := ar.handler.Group("auth")
	{
		auth.POST("/sign_in", ar.controller.SignIn)
		auth.POST("/sign_up", ar.controller.SignUp)
		auth.POST("/refresh_tokens", ar.controller.RefreshTokens)
		auth.POST("/logout", ar.controller.Logout)
	}
}
