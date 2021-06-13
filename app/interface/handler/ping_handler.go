package handler

import (
	"context"
	"io"
	"net/http"

	"go.uber.org/zap"

	"github.com/bamboooo-dev/meshi-api/app/registry"
)

type PingHandler struct {
	logger *zap.SugaredLogger
}

func NewPingHandler(l *zap.SugaredLogger, r registry.Registry) PingHandler {
	return PingHandler{
		logger: l,
	}
}

func (h PingHandler) Handle(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "pong")
		if err != nil {
			h.logger.Error(ctx, err)
			return
		}
	}
}
