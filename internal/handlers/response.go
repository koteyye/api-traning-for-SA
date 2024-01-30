package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	ctApplicationJSON = "application/json"
)

var (
	dolboebErr = errors.New("в этом методе не должно быть никакого запроса, твое видео: https://youtu.be/zdzU9R-aSqQ?si=Tvs0QUd1kaAxxz-2")
)

type errorJSON struct {
	Message string `json:"msg"`
}

func mapResponseDolboeb(w http.ResponseWriter, code int, result string) {
	rawResponse, err := json.Marshal(&errorJSON{Message: result})
	if err != nil {
		mapAltResponse(w, http.StatusInternalServerError, err.Error())
	}
	w.Header().Add("Content-Type", ctApplicationJSON)
	w.WriteHeader(code)
	w.Write(rawResponse)
}

func mapAltResponse(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(msg))
}