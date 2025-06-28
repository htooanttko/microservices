package models

import (
	"time"

	"github.com/htooanttko/microservices/services/products/internal/database"
)

type Product struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func DatabaseProductToProduct(p database.Product) Product {
	return Product{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Name:      p.Name,
	}
}

func DatabaseProductsToProducts(ps []database.Product) []Product {
	s := make([]Product, len(ps))
	for i, p := range ps {
		s[i] = DatabaseProductToProduct(p)
	}
	return s
}
