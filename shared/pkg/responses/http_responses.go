package responses

import (
	"net/http"

	"github.com/htooanttko/microservices/shared/pkg/logger"
	"github.com/htooanttko/microservices/shared/pkg/utils"
)

func WithJSON(rw http.ResponseWriter, c int, i interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(c)

	err := utils.ToJson(i, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func WithError(rw http.ResponseWriter, c int, m string) {
	if c > 499 {
		logger.Error.Printf("Responding with %v error: %v", c, m)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	WithJSON(rw, c, errorResponse{
		Error: m,
	})
}
