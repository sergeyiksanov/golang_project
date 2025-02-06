package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sergeyiksanov/golang_project/gateway/internal/lib"
	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"passwotd"`
}

func (s *SignInRequest) ToCredentialsModel() *models.Credentials {
	return &models.Credentials{
		Email:    s.Email,
		Password: s.Password,
	}
}

type SignInResponse struct {
	Access  string `json:"access,omitempty"`
	Refresh string `json:"refresh,omitempty"`
}

func SignInResponseFromTokensModel(m *models.Tokens) *SignInResponse {
	return &SignInResponse{
		Access:  m.Access,
		Refresh: m.Refresh,
	}
}

func (a *AuthController) SignIn(c *gin.Context) {
	var request SignInRequest
	if err := c.ShouldBindBodyWithJSON(request); err != nil {
		lib.AbortErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	tokens, err := a.useCase.SignIn(c.Request.Context(), request.ToCredentialsModel())
	if err != nil {
		lib.AbortErrorResponse(c, http.StatusInternalServerError, "Internal server error") // TODO: ("Valid errors")
		return
	}

	c.JSON(http.StatusOK, SignInResponseFromTokensModel(tokens))
}
