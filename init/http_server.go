package init

import (
	"isonetric-mmo-backend/internal/model"
	"isonetric-mmo-backend/internal/transport/web"
	"net/http"
	"strconv"
)

func HttpServer(config *model.ServerConfig, handler http.Handler) (*web.HttpServer, error) {
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Port),
		Handler: handler,
	}

	return web.NewHttpServer(server), nil
}
