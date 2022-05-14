package httphandler

import (
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	mux    *http.ServeMux
	logger *zap.SugaredLogger
}

func New(s *http.ServeMux, logger *zap.SugaredLogger) *Handler {
	h := Handler{s, logger}
	h.registerRoutes()

	return &h
}

func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/", h.hello)
}

func (h *Handler) hello(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("test logs")
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))

}
