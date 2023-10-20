package main

import (
	"context"
	"frame/internal/api/app"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err)
	}

	db := a.DB()
	defer db.Close()

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}
