package app

import (
	"context"

	"github.com/dimoktorr/clean_architecture/internal/app"
	"github.com/dimoktorr/clean_architecture/internal/pkg/config"
)

type ShopService struct {
	service *app.Service
}

func NewShopService() *ShopService {
	return &ShopService{}
}

func (a *ShopService) Init(ctx context.Context, cfg *config.Config) error {
	//инициализация grpc, http, роутинг, адаптеров, репозиториев, кафка, коннекторов к другим микросервисам,

	a.service = app.NewService()

	return nil
}

func (a *ShopService) Start(ctx context.Context) error {
	//старт серверов grpc, http

	return nil
}

func (a *ShopService) Stop(ctx context.Context) error {
	<-ctx.Done()

	// остановка приложения, gracefully shutdown
	return nil
}
