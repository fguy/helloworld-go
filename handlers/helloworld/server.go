package helloworld

import (
	"net/http"

	"github.com/fguy/helloworld-go/config"

	"go.uber.org/zap"
)

// NewHTTPServer -
func NewHTTPServer(cfg *config.AppConfig, logger *zap.Logger, handler *Handler) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
}
