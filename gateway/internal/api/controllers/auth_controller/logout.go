package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sergeyiksanov/golang_project/gateway/internal/lib"
	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

type LogoutRequest struct {
	Access  string `json:"access,omitempty"`
	Refresh string `json:"refresh,omitempty"`
}

func (s *LogoutRequest) toTokens() *models.Tokens {
	return &models.Tokens{
		Access:  s.Access,
		Refresh: s.Refresh,
	}
}

func (a *AuthController) Logout(c *gin.Context) {
	var request *LogoutRequest
	if err := c.ShouldBindBodyWithJSON(request); err != nil {
		lib.AbortErrorResponse(c, http.StatusBadRequest, "Ivalid request body")
		return
	}

	if err := a.useCase.Logout(c.Request.Context(), request.toTokens()); err != nil {
		lib.AbortErrorResponse(c, http.StatusInternalServerError, "Interanal server error")
		return
	}

	c.JSON(http.StatusOK, nil)
}
