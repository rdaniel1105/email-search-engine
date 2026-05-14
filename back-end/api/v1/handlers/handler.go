package handlers

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rdaniel1105/email-search-engine/back-end/helpers"
	"github.com/rdaniel1105/email-search-engine/back-end/models"
)

const maxBodyBytes = 64 * 1024 // 64 KB cap for incoming search-request bodies

var (
	errReadingRequestBody = errors.New("could not read request body")
	errInternal           = errors.New("internal server error")
)

// Routes creates a route for searching data in the API.
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", ListEmails)

	return r
}

// ListEmails handles the email search request, validates input, queries
// ZincSearch through the helpers layer, and writes the result as JSON.
func ListEmails(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("read request body: %v", err)
		helpers.WriteJSONError(w, http.StatusBadRequest, errReadingRequestBody)
		return
	}

	term, err := helpers.ValidateBody(requestBody)
	if err != nil {
		helpers.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	query, err := helpers.GetQueryParamsForRequest(r, term)
	if err != nil {
		helpers.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	result, err := helpers.DoRequest(query)
	if err != nil {
		log.Printf("DoRequest: %v", err)
		helpers.WriteJSONError(w, http.StatusInternalServerError, errInternal)
		return
	}

	if err := helpers.JSONResponse(w, http.StatusOK, models.ToSearchResponse(result)); err != nil {
		log.Printf("write search response: %v", err)
	}
}
