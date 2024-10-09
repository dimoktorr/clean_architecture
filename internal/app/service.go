package app

import (
	"context"
	"github.com/dimoktorr/clean_architecture/pkg/models"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) Get(ctx context.Context, id int32) (*models.Product, error) {
	return s.repo.Get(ctx, id)
}
