package auth_adapter

import (
	"context"

	auth_proto "github.com/sergeyiksanov/golang_project/auth_service/pkg/api/v1"
	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

func (a *AuthAdapter) SignUp(ctx context.Context, credentials *models.Credentials) error {
	request := auth_proto.SignUpRequest{
		Credentials: &auth_proto.Credentials{
			Email:    credentials.Email,
			Password: credentials.Password,
		},
	}

	_, err := a.grpcClient.SignUp(ctx, &request)
	if err != nil {
		return err
	}

	return nil
}
