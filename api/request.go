package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func urlParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}
