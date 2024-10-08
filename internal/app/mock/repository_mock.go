package mock

import (
	"context"
	"github.com/dimoktorr/clean_architecture/pkg/models"
)

type ProductsRepositoryMock struct {
	Product *models.Product
}

func (p ProductsRepositoryMock) Get(ctx context.Context, id int32) (*models.Product, error) {
	return p.Product, nil
}

func (p ProductsRepositoryMock) List(ctx context.Context) (models.Products, error) {
	//TODO implement me
	panic("implement me")
}
