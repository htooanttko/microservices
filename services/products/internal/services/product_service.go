package services

import (
	"context"
	"time"

	"github.com/htooanttko/microservices/services/products/internal/database"
	"github.com/htooanttko/microservices/services/products/internal/repositories"
)

type ProductService interface {
	GetAllProducts(ctx context.Context) ([]database.Product, error)
	GetProductByID(ctx context.Context, id int32) (database.Product, error)
	CreateProduct(ctx context.Context, p database.Product) (database.Product, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (ps *productService) GetAllProducts(ctx context.Context) ([]database.Product, error) {
	return ps.repo.GetAll(ctx)
}

func (ps *productService) GetProductByID(ctx context.Context, id int32) (database.Product, error) {
	return ps.repo.GetByID(ctx, id)
}

func (ps *productService) CreateProduct(ctx context.Context, p database.Product) (database.Product, error) {
	return ps.repo.Create(ctx, database.CreateProductParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      p.Name,
	})
}
