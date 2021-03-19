package main

import (
	"flag"
	"github.com/matthewdaltonamount/collector/pkg/heartbeat"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"

	"github.com/matthewdaltonamount/collector/pkg/aggregator"
)

func main() {
	flag.Parse()

	r := setupRouter()

	logrus.Fatal(http.ListenAndServe(":3000", r))

}

func setupRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("welcome"))
	})

	mount(r, "/collector", aggregator.Router)
	mount(r, "/heartbeat", heartbeat.Router)

	return r
}

func mount(r chi.Router, path string, routerMaker func() (http.Handler, error)) {
	handler, err := routerMaker()
	if err != nil {
		logrus.Fatal(err)
	}
	r.Mount(path, handler)
}
