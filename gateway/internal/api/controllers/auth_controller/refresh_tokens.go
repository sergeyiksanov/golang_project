package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sergeyiksanov/golang_project/gateway/internal/lib"
	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

type RefreshTokensRequest struct {
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (s *RefreshTokensRequest) toRefreshToken() models.RefreshToken {
	return models.RefreshToken(s.RefreshToken)
}

type RefreshTokensResponse struct {
	Access  string `json:"access,omitempty"`
	Refresh string `json:"refresh,omitempty"`
}

func refreshTokensResponseFromTokens(m *models.Tokens) *RefreshTokensResponse {
	return &RefreshTokensResponse{
		Access:  m.Access,
		Refresh: m.Refresh,
	}
}

func (a *AuthController) RefreshTokens(c *gin.Context) {
	var request *RefreshTokensRequest
	if err := c.ShouldBindBodyWithJSON(request); err != nil {
		lib.AbortErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	tokens, err := a.useCase.RefreshTokens(c.Request.Context(), request.toRefreshToken())
	if err != nil {
		lib.AbortErrorResponse(c, http.StatusInternalServerError, "Internal server error") //TODO: Validate err
		return
	}

	c.JSON(http.StatusOK, refreshTokensResponseFromTokens(tokens))
}
