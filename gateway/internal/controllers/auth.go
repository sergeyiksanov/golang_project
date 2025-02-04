package controllers

import (
	"net/http"

	"github.com/sergeyiksanov/golang_project/internal/controllers/requests"
	"github.com/sergeyiksanov/golang_project/internal/controllers/responses"
	"github.com/sergeyiksanov/golang_project/internal/usecases"
	"github.com/sergeyiksanov/golang_project/internal/usecases/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	logger  *zap.Logger
	useCase usecases.IAuthUseCase
}

// Logout implements IAuthController.
func (a *AuthController) Logout(c *gin.Context) {
	panic("unimplemented")
}

// RefreshTokens implements IAuthController.
func (a *AuthController) RefreshTokens(c *gin.Context) {
	panic("unimplemented")
}

// SignIn implements IAuthController.
func (a *AuthController) SignIn(c *gin.Context) {
	panic("unimplemented")
}

// SignUp implements IAuthController.
func (a *AuthController) SignUp(c *gin.Context) {
	var request requests.SignUpRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		responses.AbortErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	credentials := models.Credentials{
		Email:    request.Email,
		Passwrod: request.Passwotd,
	}
	tokens, err := a.useCase.SignUp(&credentials)
	if err != nil {
		responses.AbortErrorResponse(c, http.StatusInternalServerError, "Internal server error") // TODO: ("Valid errors")
	}

	c.JSON(http.StatusOK, responses.TokensResponse{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	})
}

func NewAuthController(logger *zap.Logger, usecase usecases.IAuthUseCase) IAuthController {
	return &AuthController{
		logger:  logger,
		useCase: usecase,
	}
}
