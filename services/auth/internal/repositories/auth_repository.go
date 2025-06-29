package repositories

import (
	"context"
	"database/sql"

	"github.com/htooanttko/microservices/services/auth/internal/database"
)

type AuthRepository interface {
	Create(ctx context.Context, arg database.CreateUserParams) (database.User, error)
	GetByEmail(ctx context.Context, em string) (database.User, error)
}

type authRepo struct {
	q *database.Queries
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepo{q: database.New(db)}
}

func (ar *authRepo) Create(ctx context.Context, arg database.CreateUserParams) (database.User, error) {
	return ar.q.CreateUser(ctx, arg)
}

func (ar *authRepo) GetByEmail(ctx context.Context, em string) (database.User, error) {
	return ar.q.GetUserByEmail(ctx, em)
}
