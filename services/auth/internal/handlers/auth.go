package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/htooanttko/microservices/services/auth/internal/database"
	"github.com/htooanttko/microservices/services/auth/internal/models"
	"github.com/htooanttko/microservices/services/auth/internal/services"
	"github.com/htooanttko/microservices/shared/pkg/responses"
	"github.com/htooanttko/microservices/shared/pkg/utils"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

var validate = validator.New()

func (ah *AuthHandler) SignUp(rw http.ResponseWriter, r *http.Request) {
	var du models.SignUp
	if err := utils.FromJson(&du, r.Body); err != nil {
		responses.WithError(rw, http.StatusBadRequest, err.Error())
		return
	}

	if err := validate.Struct(du); err != nil {
		responses.WithError(rw, http.StatusBadRequest, err.Error())
		return
	}

	usr, a, rf, err := ah.service.SignUpUser(r.Context(), database.User{
		Name:     du.Name,
		Email:    du.Email,
		Password: du.Password,
	})
	if err != nil {
		responses.WithError(rw, http.StatusBadRequest, err.Error())
		return
	}

	responses.WithJSON(rw, http.StatusCreated, map[string]any{
		"access_token":  a,
		"refresh_token": rf,
		"user":          models.DatabaseUserToUser(usr),
	})
}

func (ah *AuthHandler) Login(rw http.ResponseWriter, r *http.Request) {
	var du models.Login
	if err := utils.FromJson(&du, r.Body); err != nil {
		responses.WithError(rw, http.StatusBadRequest, err.Error())
		return
	}

	if err := validate.Struct(du); err != nil {
		responses.WithError(rw, http.StatusBadRequest, err.Error())
		return
	}

	usr, a, rf, err := ah.service.LoginUser(r.Context(), du.Email, du.Password)
	if err != nil {
		responses.WithError(rw, http.StatusUnauthorized, err.Error())
		return
	}

	responses.WithJSON(rw, http.StatusOK, map[string]any{
		"access_token":  a,
		"refresh_token": rf,
		"user":          models.DatabaseUserToUser(usr),
	})
}

func (ah *AuthHandler) Refresh(rw http.ResponseWriter, r *http.Request) {
	var dat map[string]string
	if err := utils.FromJson(&dat, r.Body); err != nil {
		responses.WithError(rw, http.StatusBadRequest, err.Error())
		return
	}

	a, rf, err := ah.service.RefreshToken(dat["refresh_token"])
	if err != nil {
		responses.WithError(rw, http.StatusUnauthorized, err.Error())
		return
	}

	responses.WithJSON(rw, http.StatusOK, map[string]string{
		"access_token":  a,
		"refresh_token": rf,
	})
}
