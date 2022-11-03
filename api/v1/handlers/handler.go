package handlers

import (
	"example/mamuro/helpers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Routes creates a route for searching data.
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", List)
	r.Route("/{input}", func(r chi.Router) {
		r.With(ValidateURL).Get("/", GetData())
	})

	return r
}

// ValidateURL validates the URL for the parameter route
func ValidateURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		inputParam := chi.URLParam(r, "input")
		if inputParam == "" {
			w.WriteHeader(404)
			_, err := w.Write([]byte(http.StatusText(404)))
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

// List displays the list of emails obtained from the request to ZincSearch.
func List(w http.ResponseWriter, r *http.Request) {
	query := `{
        "search_type": "alldocuments",
        "query":
        {"term": ""},
        "_source": ["From","To","Subject", "Body","Date","Message-ID"]
    }`
	helpers.DoRequest(query, w)
}

// GetData makes a request to ZincSearch to get the data requested.
func GetData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		term := chi.URLParam(r, "input")

		//validate term with Regex
		query := `{
        "search_type": "match",
        "query":
        {
            "term": "` + term + `"
        },
        "_source": ["From","To","Subject", "Body","Date","Message-ID"]
    	}`
		helpers.DoRequest(query, w)
	}
}
