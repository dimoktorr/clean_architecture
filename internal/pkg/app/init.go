package app

import (
	"context"

	"github.com/dimoktorr/clean_architecture/internal/app"
	"github.com/dimoktorr/clean_architecture/internal/pkg/config"
)

type App struct {
	service *app.Service
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init(ctx context.Context, cfg *config.Config) error {
	//инициализация grpc, http, роутинг, адаптеров, репозиториев, кафка, коннекторов к другим микросервисам,

	a.service = app.NewService()

	return nil
}

func (a *App) Start(ctx context.Context) error {
	//старт серверов grpc, http

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	<-ctx.Done()

	// остановка приложения, gracefully shutdown
	return nil
}
