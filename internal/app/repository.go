//go:generate mockgen -destination=./mock/repository_mock_generated.go -package=mock  -source repository.go
package app

import (
	"context"
	"github.com/dimoktorr/clean_architecture/pkg/models"
)

type Repository interface {
	Get(ctx context.Context, id int32) (*models.Product, error)
	List(ctx context.Context) (models.Products, error)
}
