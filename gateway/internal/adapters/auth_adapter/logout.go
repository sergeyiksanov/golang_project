package auth_adapter

import (
	"context"

	auth_proto "github.com/sergeyiksanov/golang_project/auth_service/pkg/api/v1"
	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/models"
)

func (a *AuthAdapter) Logout(ctx context.Context, t *models.Tokens) error {
	request := auth_proto.LogoutRequest{
		Tokens: &auth_proto.Tokens{
			Access:  t.Access,
			Refresh: t.Refresh,
		},
	}

	_, err := a.grpcClient.Logout(ctx, &request)
	if err != nil {
		return err
	}

	return nil
}
