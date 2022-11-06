package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx := context.Background()
	run(ctx, os.Args)
}

func run(ctx context.Context, _ []string) error {

	// CONSTRUCT CONFIGURATION
	/*var cfg struct {

	}*/

	// START DATABASE

	// INITIALIZE HTTP CLIENTS

	// ENVIRONMENT SPECIFIC LOGIC

	// Start API Service

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)

	return nil
}
