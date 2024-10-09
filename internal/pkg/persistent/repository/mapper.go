package repository

import "github.com/dimoktorr/clean_architecture/pkg/models"

func toModelsProduct(product Product) *models.Product {
	return &models.Product{
		ID: product.ID.Int32,
	}
}
