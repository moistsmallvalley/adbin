package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func withURLParams(r *http.Request, params map[string]string) *http.Request {
	rctx := chi.NewRouteContext()
	for key, value := range params {
		rctx.URLParams.Add(key, value)
	}
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}
