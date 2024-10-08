package main

import (
	"context"
	logBase "log"

	"github.com/dimoktorr/clean_architecture/internal/app/shop"
	"github.com/dimoktorr/clean_architecture/internal/app/shop/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logBase.Fatalln(err)
	}

	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := shop.NewApp()

	if err := app.Init(ctxWithCancel, cfg); err != nil {
		logBase.Fatal("start app failed")
	}

	if err := app.Start(ctxWithCancel); err != nil {
		logBase.Fatal("start app failed")
	}

	if err := app.Stop(ctxWithCancel); err != nil {
		logBase.Fatalf("stop app failed")
	}
}
