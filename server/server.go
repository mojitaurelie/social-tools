package server

import (
	"fmt"
	"log"
	"net/http"
	"social-tools/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Serve() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(recovery)
	router.Route("/api", func(rApi chi.Router) {
		rApi.Route("/v1", func(r chi.Router) {

		})
	})
	log.Printf("INFO http server started at %d", config.Current().Server.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Current().Server.Port), router)
	if err != nil {
		log.Fatal(err)
	}
}

func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				internalServerError(w, r)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
