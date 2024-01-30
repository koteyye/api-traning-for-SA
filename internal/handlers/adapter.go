package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type emptyBody struct {}

func mapRequestNoBody(r *http.Request) error {
	var body emptyBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		} else {
			return dolboebErr
		}
	}
	return dolboebErr
}