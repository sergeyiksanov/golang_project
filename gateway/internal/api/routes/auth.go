package routes

import (
	"github.com/sergeyiksanov/golang_project/internal/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthRoutes struct {
	logger     *zap.Logger
	handler    *gin.Engine
	controller controllers.IAuthController
}

func NewAuthRoutes(logger *zap.Logger, handler *gin.Engine, controller controllers.IAuthController) *AuthRoutes {
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
