package app

import (
	"github.com/sergeyiksanov/golang_project/gateway/internal/adapters/auth_adapter"
	"github.com/sergeyiksanov/golang_project/gateway/internal/api/controllers/auth_controller"
	"github.com/sergeyiksanov/golang_project/gateway/internal/api/middleware"
	"github.com/sergeyiksanov/golang_project/gateway/internal/api/routes"
	"github.com/sergeyiksanov/golang_project/gateway/internal/app/config"
	"github.com/sergeyiksanov/golang_project/gateway/internal/lib/logger"
	"github.com/sergeyiksanov/golang_project/gateway/internal/usecases/auth_usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type serviceProvider struct {
	//cfg
	cfg *config.Config

	//logger
	logger *zap.Logger

	//routes
	routes     *routes.Routes
	authRoutes *routes.AuthRoutes

	//gin
	gin *gin.Engine

	//middleware
	loggerMiddleware *middleware.LoggerMiddleware

	//controllers
	authController *auth_controller.AuthController

	//usecases
	authUseCase *auth_usecase.AuthUseCase

	//adapters
	authAdapter *auth_adapter.AuthAdapter
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *config.Config {
	if s.cfg == nil {
		cfg, err := config.Load()
		if err != nil {
			panic(err)
		}
		s.cfg = cfg
	}

	return s.cfg
}

func (s *serviceProvider) Logger() *zap.Logger {
	if s.logger == nil {
		logger, err := logger.NewLogger(&s.Config().Logger)
		if err != nil {
			return nil
		}

		s.logger = logger
	}

	return s.logger
}

func (s *serviceProvider) Routes() *routes.Routes {
	if s.routes == nil {
		s.routes = routes.NewRoutes(*s.AuthRoutes())
	}

	return s.routes
}

func (s *serviceProvider) AuthRoutes() *routes.AuthRoutes {
	if s.authRoutes == nil {
		s.authRoutes = routes.NewAuthRoutes(s.Logger(), s.GinEngine(), s.AuthController())
	}

	return s.authRoutes
}

func (s *serviceProvider) GinEngine() *gin.Engine {
	if s.gin == nil {
		gin.SetMode(gin.ReleaseMode)
		s.gin = gin.New()
		s.gin.Use(
			gin.Recovery(),
			s.LoggerMiddleware().Handler(),
		)
	}

	return s.gin
}

func (s *serviceProvider) LoggerMiddleware() *middleware.LoggerMiddleware {
	if s.loggerMiddleware == nil {
		s.loggerMiddleware = middleware.NewLoggerMiddlewate(s.Logger())
	}

	return s.loggerMiddleware
}

func (s *serviceProvider) AuthController() *auth_controller.AuthController {
	if s.authController == nil {
		s.authController = auth_controller.NewAuthController(s.Logger(), s.AuthUseCase())
	}

	return s.authController
}

func (s *serviceProvider) AuthUseCase() *auth_usecase.AuthUseCase {
	if s.authUseCase == nil {
		s.authUseCase = auth_usecase.NewAuthUseCase(s.Logger(), s.AuthAdapter())
	}

	return s.authUseCase
}

func (s *serviceProvider) AuthAdapter() *auth_adapter.AuthAdapter {
	if s.authAdapter == nil {
		s.authAdapter = auth_adapter.NewAuthAdapter(s.Logger(), &s.Config().Grpc)
	}

	return s.authAdapter
}
