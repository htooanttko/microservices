package services

import (
	"context"
	"time"

	"github.com/htooanttko/microservices/services/auth/internal/database"
	"github.com/htooanttko/microservices/services/auth/internal/repositories"
	"github.com/htooanttko/microservices/services/auth/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUpUser(ctx context.Context, u database.User) (database.User, string, string, error)
	LoginUser(ctx context.Context, em, pw string) (database.User, string, string, error)
	RefreshToken(t string) (string, string, error)
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (as *authService) SignUpUser(ctx context.Context, u database.User) (database.User, string, string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return database.User{}, "", "", err
	}

	a, r, err := utils.GenerateTokens(u.Email)
	if err != nil {
		return database.User{}, "", "", err
	}

	usr, err := as.repo.Create(ctx, database.CreateUserParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  string(hp),
	})
	if err != nil {
		return database.User{}, "", "", err
	}

	return usr, a, r, nil
}

func (as *authService) LoginUser(ctx context.Context, em, pw string) (database.User, string, string, error) {
	usr, err := as.repo.GetByEmail(ctx, em)
	if err != nil {
		return database.User{}, "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(pw)); err != nil {
		return database.User{}, "", "", err
	}

	a, r, err := utils.GenerateTokens(usr.Email)
	if err != nil {
		return database.User{}, "", "", err
	}

	return usr, a, r, nil
}

func (as *authService) RefreshToken(t string) (string, string, error) {
	em, err := utils.VerifyRefreshToken(t)
	if err != nil {
		return "", "", err
	}

	return utils.GenerateTokens(em)
}
