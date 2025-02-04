package main

import (
	"context"
	"log"

	"github.com/sergeyiksanov/AuthService/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
}
