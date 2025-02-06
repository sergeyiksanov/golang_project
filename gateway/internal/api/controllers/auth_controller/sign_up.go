package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sergeyiksanov/golang_project/gateway/internal/lib"
	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"passwotd"`
}

func (s *SignUpRequest) ToCredentialsModel() *models.Credentials {
	return &models.Credentials{
		Email:    s.Email,
		Password: s.Password,
	}
}

type SignUpResponse struct {
	Access  string `json:"access,omitempty"`
	Refresh string `json:"refresh,omitempty"`
}

func SignUpResponseFromTokensModel(m *models.Tokens) *SignUpResponse {
	return &SignUpResponse{
		Access:  m.Access,
		Refresh: m.Refresh,
	}
}

func (a *AuthController) SignUp(c *gin.Context) {
	var request *SignUpRequest
	if err := c.ShouldBindBodyWithJSON(request); err != nil {
		lib.AbortErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	err := a.useCase.SignUp(c.Request.Context(), request.ToCredentialsModel())
	if err != nil {
		lib.AbortErrorResponse(c, http.StatusInternalServerError, "Internal server error") // TODO: ("Valid errors")
		return
	}

	c.JSON(http.StatusOK, nil)
}
