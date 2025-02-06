package auth_adapter

import (
	"context"

	auth_proto "github.com/sergeyiksanov/golang_project/auth_service/pkg/api/v1"
	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

func (a *AuthAdapter) RefreshTokens(ctx context.Context, token models.RefreshToken) (*models.Tokens, error) {
	request := auth_proto.RefreshTokensRequest{
		RefreshToken: string(token),
	}

	response, err := a.grpcClient.RefreshTokens(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &models.Tokens{
		Access:  response.Tokens.Access,
		Refresh: response.Tokens.Refresh,
	}, nil
}
