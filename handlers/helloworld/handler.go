package helloworld

import (
	"context"
	"encoding/json"
	"net/http"

	repo "github.com/fguy/helloworld-go/repositories/helloworld"
	"go.uber.org/zap"
)

// Handler -
type Handler struct {
	logger     *zap.Logger
	repository repo.Interface
}

// ServeHTTP implements http.Handler
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	h.logger.Info("request", zap.String("title", title))
	page, err := h.repository.GetPage(context.Background(), title)
	if err != nil {
		h.logger.Error("error", zap.Error(err))
		w.WriteHeader(500)
	} else if page != nil {
		h.logger.Debug("page", zap.Any("page", page))
		json.NewEncoder(w).Encode(page)
	} else {
		w.WriteHeader(404)
	}
}

// NewHandler -
func NewHandler(logger *zap.Logger, repository repo.Interface) *Handler {
	logger.Info("executing NewHandler")
	return &Handler{
		logger:     logger,
		repository: repository,
	}
}
