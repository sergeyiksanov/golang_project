package main

import (
	"context"

	"github.com/sergeyiksanov/golang_project/gateway/internal/app"
)

func main() {
	app, err := app.NewApp(context.Background())
	if err != nil {
		panic(err)
	}

	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
