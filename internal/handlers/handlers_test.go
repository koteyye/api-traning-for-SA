package handlers

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlers_NewHandlers(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		handlers := NewHandlers(slog.Logger{})
		assert.Equal(t, &Handlers{
			logger: slog.Logger{},
		}, handlers)
	})
}

func testInitHandler(t *testing.T) (*Handlers) {
	handler := NewHandlers(slog.Logger{})
	return handler
}

func TestHandlers_Health(t *testing.T) {
	t.Run("health", func(t *testing.T) {
		t.Run("ok", func(t *testing.T) {
			h := testInitHandler(t)
			r := httptest.NewRequest(http.MethodGet, "/api/traning/health", nil)
			w := httptest.NewRecorder()
			h.Health(w, r)

			assert.Equal(t, http.StatusOK, w.Code)
		})
		t.Run("haveBody", func(t *testing.T) {
			testRequest := `{"testField": 1}`

			h := testInitHandler(t)
			r := httptest.NewRequest(http.MethodGet, "/api/traning/health", strings.NewReader(testRequest))
			w := httptest.NewRecorder()
			h.Health(w, r)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, ctApplicationJSON, w.Header().Get("Content-Type"))
		})
	})
}