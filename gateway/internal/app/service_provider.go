package app

import (
	"github.com/sergeyiksanov/golang_project/internal/api/middleware"
	"github.com/sergeyiksanov/golang_project/internal/api/routes"
	"github.com/sergeyiksanov/golang_project/internal/app/logger"
	"github.com/sergeyiksanov/golang_project/internal/clients"
	"github.com/sergeyiksanov/golang_project/internal/config"
	"github.com/sergeyiksanov/golang_project/internal/controllers"
	"github.com/sergeyiksanov/golang_project/internal/usecases"

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
	authController controllers.IAuthController

	//usecases
	authUseCase usecases.IAuthUseCase

	//clients
	authClient clients.IAuthClient
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

func (s *serviceProvider) AuthController() controllers.IAuthController {
	if s.authController == nil {
		s.authController = controllers.NewAuthController(s.Logger(), s.AuthUseCase())
	}

	return s.authController
}

func (s *serviceProvider) AuthUseCase() usecases.IAuthUseCase {
	if s.authUseCase == nil {
		s.authUseCase = usecases.NewAuthUseCase(s.Logger(), s.AuthClient())
	}

	return s.authUseCase
}

func (s *serviceProvider) AuthClient() clients.IAuthClient {
	if s.authClient == nil {
		s.authClient = clients.NewAuthClient(s.Logger())
	}

	return s.authClient
}
