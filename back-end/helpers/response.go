package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rdaniel1105/email-search-engine/back-end/models"
)

// JSONResponse serializes payload and writes it as a JSON response.
func JSONResponse(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("json marshal response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if _, err := w.Write(response); err != nil {
		return fmt.Errorf("json write response: %w", err)
	}

	return nil
}

// WriteJSONError writes a JSON error response with the given status code.
// The err.Error() is exposed to the client, so callers must only pass errors
// safe for external consumption — never wrap internal/DB errors here.
func WriteJSONError(w http.ResponseWriter, code int, err error) {
	if writeErr := JSONResponse(w, code, models.ErrorResponseMessage{"message": err.Error()}); writeErr != nil {
		log.Printf("write json error: %v", writeErr)
	}
}
