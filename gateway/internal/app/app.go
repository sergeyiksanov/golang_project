package app

import (
	"context"

	"go.uber.org/zap"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	defer a.serviceProvider.Logger().Sync()

	a.serviceProvider.Logger().Info(
		"Application started",
		zap.String("version", "1.0.0"),
		zap.Int("port", 8080),
	)

	if err := a.runDeps(ctx); err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.setupRoutes,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) runDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.runGin,
		a.connectAuthAdapter,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) setupRoutes(_ context.Context) error {
	a.serviceProvider.Routes().Setup()
	return nil
}

func (a *App) runGin(_ context.Context) error {
	return a.serviceProvider.GinEngine().Run()
}

func (a *App) connectAuthAdapter(_ context.Context) error {
	return a.serviceProvider.AuthAdapter().Connect()
}
