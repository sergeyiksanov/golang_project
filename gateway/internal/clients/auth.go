package clients

import (
	"github.com/sergeyiksanov/golang_project/internal/usecases/models"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AuthClient struct {
	logger   *zap.Logger
	grpcConn *grpc.ClientConn
}

// Logout implements IAuthClient.
func (a *AuthClient) Logout() {
	panic("unimplemented")
}

// RefreshTokens implements IAuthClient.
func (a *AuthClient) RefreshTokens() {
	panic("unimplemented")
}

// SignIn implements IAuthClient.
func (a *AuthClient) SignIn() {
	panic("unimplemented")
}

// SignUp implements IAuthClient.
func (a *AuthClient) SignUp(credentials *models.Credentials) (*models.Tokens, error) {
	panic("unimplemented")
}

// VerifyAccessToken implements IAuthClient.
func (a *AuthClient) VerifyAccessToken() {
	panic("unimplemented")
}

func NewAuthClient(logger *zap.Logger) IAuthClient {
	return &AuthClient{
		logger: logger,
	}
}
