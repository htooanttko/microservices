package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/htooanttko/microservices/services/products/internal/database"
	"github.com/htooanttko/microservices/services/products/internal/models"
	"github.com/htooanttko/microservices/services/products/internal/services"
	"github.com/htooanttko/microservices/shared/pkg/responses"
	"github.com/htooanttko/microservices/shared/pkg/utils"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (ph *ProductHandler) GetAll(rw http.ResponseWriter, r *http.Request) {
	p, err := ph.service.GetAllProducts(r.Context())
	if err != nil {
		responses.WithError(rw, http.StatusInternalServerError, err.Error())
		return
	}
	responses.WithJSON(rw, http.StatusOK, models.DatabaseProductsToProducts(p))
}

func (ph *ProductHandler) GetByID(rw http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responses.WithError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	p, err := ph.service.GetProductByID(r.Context(), int32(id))
	if err != nil {
		responses.WithError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	responses.WithJSON(rw, http.StatusOK, models.DatabaseProductToProduct(p))
}

func (ph *ProductHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var dp database.Product
	if err := utils.FromJson(&dp, r.Body); err != nil {
		responses.WithError(rw, http.StatusBadRequest, err.Error())
		return
	}

	p, err := ph.service.CreateProduct(r.Context(), dp)
	if err != nil {
		responses.WithError(rw, http.StatusInternalServerError, err.Error())
	}

	responses.WithJSON(rw, http.StatusCreated, models.DatabaseProductToProduct(p))
}
