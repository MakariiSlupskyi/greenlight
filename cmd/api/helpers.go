package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
