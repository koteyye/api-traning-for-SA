package handlers

import (
	"net/http"
)

// Health godoc
// @Tags Health
// @Summary Запрос для проверки работы сервиса
// @Accept plain
// @Produce plain
// @Success 200 "Все заебись"
// @Failure 400 "Ты долбаеб"
// @Failure 500 "Все хуево"
// @Router /health [get]
func (h Handlers) Health(w http.ResponseWriter, r *http.Request) {

	err := mapRequestNoBody(r)
	if err != nil {
		mapResponseDolboeb(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}