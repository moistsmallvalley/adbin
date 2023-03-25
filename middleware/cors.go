package middleware

import "net/http"

type corsMiddleware struct {
	next        http.Handler
	allowOrigin string
	allowMethod string
}

func NewCORSMiddleware(next http.Handler, allowOrigin, allowMethod string) http.Handler {
	return &corsMiddleware{
		next:        next,
		allowOrigin: allowOrigin,
		allowMethod: allowMethod,
	}
}

func (h *corsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", h.allowOrigin)
	w.Header().Set("Access-Control-Allow-Method", h.allowMethod)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	h.next.ServeHTTP(w, r)
}
