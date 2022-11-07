package router

import (
	"example/mamuro/api/v1/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

const (
	searchPath    = "/api/search"
	allowedOrigin = "http://localhost:8080"
)

// ServeRouter serves the router in which the server will be running
func ServeRouter() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := Initialize()

	fmt.Printf("server running in port:%+v", port)
	server := http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		Handler:     r,
		ReadTimeout: 1000 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

// Initialize initializes the Server
func Initialize() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), //forces Content-type
		middleware.RedirectSlashes,
		middleware.Logger,
		middleware.Recoverer, //middleware to recover from panics
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{allowedOrigin},
			AllowedMethods:   []string{http.MethodGet},
			AllowedHeaders:   []string{"Accept", "Content-Type"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	router.Use(middleware.Timeout(30 * time.Second))

	router.Route(searchPath, func(r chi.Router) {
		r.Mount("/", handlers.Routes())
	})

	return router
}
