package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jasonlvhit/gocron"
)

func task1() {
	fmt.Println("I am running task 1.")
	gocron.Remove(task1)
}

func task2() {
	fmt.Println("I am running task 2.")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)

	// gocron.Every(1).Second().Do(task1)
	// gocron.Every(1).Second().Do(task2)
	// <-gocron.Start()
}
