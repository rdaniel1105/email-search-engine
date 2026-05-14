package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rdaniel1105/email-search-engine/back-end/models"
)

// JSONResponse serializes payload and writes it as a JSON response.
// HTML-escape is disabled so characters like '<', '>', '&' (common in
// Message-IDs and email bodies) render as themselves rather than the
// noisy < / > / & escapes.
func JSONResponse(w http.ResponseWriter, code int, payload interface{}) error {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(payload); err != nil {
		return fmt.Errorf("json marshal response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if _, err := w.Write(buf.Bytes()); err != nil {
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
