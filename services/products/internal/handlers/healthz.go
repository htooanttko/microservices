package handlers

import (
	"net/http"

	"github.com/htooanttko/microservices/shared/pkg/responses"
)

func GetHealthz(rw http.ResponseWriter, r *http.Request) {
	type res struct {
		Message string `json:"message"`
	}

	responses.WithJSON(rw, 200, res{
		Message: "success",
	})
}
