package repositories

import (
	"context"
	"database/sql"

	"github.com/htooanttko/microservices/services/products/internal/database"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]database.Product, error)
	GetByID(ctx context.Context, id int32) (database.Product, error)
	Create(ctx context.Context, arg database.CreateProductParams) (database.Product, error)
}

type productRepo struct {
	q *database.Queries
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepo{q: database.New(db)}
}

func (r *productRepo) GetAll(ctx context.Context) ([]database.Product, error) {
	return r.q.GetProducts(ctx)
}

func (r *productRepo) GetByID(ctx context.Context, id int32) (database.Product, error) {
	product, err := r.q.GetProductByID(ctx, id)
	// if errors.Is(err, sql.ErrNoRows) {
	// 	return nil, nil
	// }
	return product, err
}

func (r *productRepo) Create(ctx context.Context, arg database.CreateProductParams) (database.Product, error) {
	return r.q.CreateProduct(ctx, arg)
}
