package repository

import (
	"context"
	"github.com/dimoktorr/clean_architecture/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pgx *pgxpool.Pool
}

func NewRepository(pgx *pgxpool.Pool) *Repository {
	return &Repository{
		pgx: pgx,
	}
}

func (r Repository) Get(ctx context.Context, id int32) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) List(ctx context.Context) (models.Products, error) {
	//TODO implement me
	panic("implement me")
}
