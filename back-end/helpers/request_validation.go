package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
)

// BodyParsed represents the parsed body from the request the client made.
type BodyParsed struct {
	Term string `json:"term"`
}

var (
	errMissingBody    = errors.New("missing body")
	errMissingTermKey = errors.New(`"term" key missing`)
	errInvalidJSON    = errors.New("invalid json body")
)

// ValidateBody parses the request body and returns the search term, or an
// error suitable for surfacing to the client (no internal details leaked).
func ValidateBody(body []byte) (string, error) {
	if len(body) == 0 {
		return "", errMissingBody
	}

	var parsed BodyParsed
	if err := json.Unmarshal(body, &parsed); err != nil {
		return "", fmt.Errorf("%w: %s", errInvalidJSON, err)
	}

	if parsed.Term == "" {
		return "", errMissingTermKey
	}

	return parsed.Term, nil
}
