package auth_adapter

import (
	"fmt"

	"github.com/sergeyiksanov/golang_project/gateway/internal/app/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	auth_proto "github.com/sergeyiksanov/golang_project/auth_service/pkg/api/v1"
)

type AuthAdapter struct {
	logger     *zap.Logger
	cfg        *config.GrpcConfig
	grpcClient auth_proto.AuthClient
}

func (a *AuthAdapter) Connect() error {
	conn, err := grpc.NewClient(a.cfg.ServerAddress)
	if err != nil {
		return fmt.Errorf("connect to auth server error: %w", err)
	}

	client := auth_proto.NewAuthClient(conn)
	a.grpcClient = client

	return nil
}

func NewAuthAdapter(logger *zap.Logger, cfg *config.GrpcConfig) *AuthAdapter {
	return &AuthAdapter{
		logger: logger,
		cfg:    cfg,
	}
}
