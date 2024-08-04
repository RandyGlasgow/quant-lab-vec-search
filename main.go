package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type GlobalRepository struct {
	// db for postgres connection
	PgDB *sql.DB
}

func main() {
	// create a new chi router
	r := chi.NewRouter()

	// create a new logger
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the chi router"))
	})

	fmt.Println("Server is running on port 8080")
	// serve the router on port 8080
	http.ListenAndServe(":8080", r)
}
