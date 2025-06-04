package initialization

import (
	"isonetric-mmo-backend/pkg/model"
	"net/http"
	"strconv"
)

func Server(config *model.ServerConfig, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":" + strconv.Itoa(config.Port),
		Handler: handler,
	}
}
