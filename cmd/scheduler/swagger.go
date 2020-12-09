package main

import (
	"log"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

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
