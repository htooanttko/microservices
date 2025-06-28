-- name: CreateProduct :one
INSERT INTO products (created_at, updated_at, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProductByID :one
SELECT * FROM products WHERE id = $1;