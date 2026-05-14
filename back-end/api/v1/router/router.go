package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/rdaniel1105/email-search-engine/back-end/api/v1/handlers"
)

const (
	searchPath    = "/api/v1/emails/search"
	allowedOrigin = "http://localhost:8080"
	defaultPort   = "3000"

	readTimeout    = 5 * time.Second
	writeTimeout   = 30 * time.Second
	requestTimeout = 30 * time.Second
	corsMaxAge     = 300
)

// ServeRouter starts the HTTP server.
func ServeRouter() {
	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file loaded: %v (falling back to environment)", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := Initialize()

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("server listening on :%s", port)
	log.Fatal(server.ListenAndServe())
}

// Initialize wires up middleware and routes and returns the chi router.
func Initialize() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler(cors.Options{
			AllowedOrigins: []string{allowedOrigin},
			AllowedMethods: []string{http.MethodPost},
			AllowedHeaders: []string{
				"Accept",
				"Content-Type",
				"Access-Control-Allow-Origin",
				"Access-Control-Allow-Credentials",
			},
			AllowCredentials: false,
			MaxAge:           corsMaxAge,
		}),
	)

	router.Use(middleware.Timeout(requestTimeout))

	router.Route(searchPath, func(r chi.Router) {
		r.Mount("/", handlers.Routes())
	})

	return router
}
