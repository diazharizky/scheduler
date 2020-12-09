package main

import (
	"log"
	"net/http"

	"github.com/go-chi/cors"
	"github.com/gobuffalo/packr/v2"
)

func getAllowAllCORS() func(http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler
}

func swaggerHandler() (http.HandlerFunc, error) {
	box := packr.New("api", "../../api/swagger-spec/")
	swaggerSource, err := box.FindString("scheduler.json")
	if err != nil {
		log.Fatal(err)
	}

	sh := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		// nolint:errcheck
		w.Write([]byte(swaggerSource))
	}

	return http.HandlerFunc(sh), nil
}
