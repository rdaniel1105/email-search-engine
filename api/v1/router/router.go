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

// ServeRouter serves the router in which the server will be running
func ServeRouter() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	r := Initialize()

	fmt.Printf("Server running in port:%+v", port)
	server := http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		Handler:     r,
		ReadTimeout: 1000 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

// Initialize initializes the Server
func Initialize() *chi.Mux {
	searchPath := "/api/search"

	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), //forces Content-type
		middleware.RedirectSlashes,
		middleware.Logger,
		middleware.Recoverer, //middleware to recover from panics
		//middleware.Heartbeat("/health"), //for heartbeat process such as Kubernetes liveprobeness
		cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
	)

	//Sets context for all requests
	router.Use(middleware.Timeout(30 * time.Second))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome"))
		if err != nil {
			fmt.Println(err)
		}
	})

	//routes
	router.Route(searchPath, func(r chi.Router) {
		r.Mount("/", handlers.Routes())
	})

	return router
}
