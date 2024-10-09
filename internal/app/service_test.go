package app

import (
	"context"
	"github.com/dimoktorr/clean_architecture/internal/app/mock"
	"github.com/dimoktorr/clean_architecture/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProduct(t *testing.T) {
	//arrange

	repo := mock.ProductsRepositoryMock{
		Product: &models.Product{
			ID: 4,
		},
	}

	service := NewService(repo)

	want := &models.Product{
		ID: 4,
	}

	//act
	got, err := service.Get(context.Background(), 4)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
