package aggregator

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"
)

func Router() (http.Handler, error) {
	r := chi.NewRouter()
	r.Use(logCtx)
	r.Get("/", parseLogService)

	return r, nil
}

func logCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := "/log.txt"
		ctx = context.WithValue(ctx, "log", log)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
