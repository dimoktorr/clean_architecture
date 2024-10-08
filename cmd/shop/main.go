package main

import (
	"context"
	logBase "log"

	"github.com/dimoktorr/clean_architecture/internal/pkg/app"
	"github.com/dimoktorr/clean_architecture/internal/pkg/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logBase.Fatalln(err)
	}

	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := app.NewShopService()

	if err := service.Init(ctxWithCancel, cfg); err != nil {
		logBase.Fatal("start app failed")
	}

	if err := service.Start(ctxWithCancel); err != nil {
		logBase.Fatal("start app failed")
	}

	if err := service.Stop(ctxWithCancel); err != nil {
		logBase.Fatalf("stop app failed")
	}
}
