package handlers

import (
	"log/slog"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/go-chi/chi"

	_ "github.com/koteyye/api-traning-for-SA/docs"
)

// Handlers http обработчик
type Handlers struct {
	logger slog.Logger
}

// NewHandlers возвращает новый экземпляр http обработчика
func NewHandlers(logger slog.Logger) *Handlers {
	return &Handlers{logger: logger}
}

// InitRoutes инициализация роутов
func (h *Handlers) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("swagger/doc.json"),  //The url pointing to API definition
	))

	r.Route("/api/traning", func(r chi.Router) {
		r.Get("/health", h.Health)
	})
	return r
}

