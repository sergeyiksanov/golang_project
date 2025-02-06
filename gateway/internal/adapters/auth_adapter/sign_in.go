package auth_adapter

import (
	"context"

	auth_proto "github.com/sergeyiksanov/golang_project/auth_service/pkg/api/v1"
	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

func (a *AuthAdapter) SignIn(ctx context.Context, credentials *models.Credentials) (*models.Tokens, error) {
	request := auth_proto.SignInRequest{
		Credentials: &auth_proto.Credentials{
			Email:    credentials.Email,
			Password: credentials.Password,
		},
	}

	response, err := a.grpcClient.SignIn(ctx, &request)
	if err != nil {
		return nil, err
	}

	tokens := models.Tokens{
		Access:  response.Tokens.Access,
		Refresh: response.Tokens.Refresh,
	}
	return &tokens, nil
}
