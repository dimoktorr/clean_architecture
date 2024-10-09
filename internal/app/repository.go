package app

import (
	"context"
	"github.com/dimoktorr/clean_architecture/pkg/models"
)

type Repository interface {
	Get(ctx context.Context, id int32) (*models.Product, error)
	List(ctx context.Context) (models.Products, error)
}
