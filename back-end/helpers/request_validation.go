package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// BodyParsed represents the parsed body from the request the client made.
type BodyParsed struct {
	Term string `json:"term"`
}

var (
	bodyParsed BodyParsed

	errMissingBody    = errors.New("missing body")
	errMissingTermKey = errors.New(`"term" key missing`)
	errIsNaN          = errors.New(`is not a number`)
)

// ValidateBody parses the body from the request, verifies the body is not empty
// and checks the key "term" comes in the body; then returns the validated body.
func ValidateBody(w http.ResponseWriter, body []byte) (string, error) {
	err := json.Unmarshal(body, &bodyParsed)
	if err != nil {
		return "", fmt.Errorf("response writer: %w", err)
	}

	if string(body) == "" {
		return "", ResponseErrorHelper(w, http.StatusBadRequest, errMissingBody.Error(), errMissingBody)
	}

	if !strings.Contains(string(body), "term") {
		return "", ResponseErrorHelper(w, http.StatusBadRequest, errMissingTermKey.Error(), errMissingTermKey)
	}

	return bodyParsed.Term, nil
}

// ValidateQueryParams validates that the query params are numbers.
func ValidateQueryParams(queryParam string) error {
	f, err := strconv.ParseFloat(queryParam, 64)
	if err != nil {
		return fmt.Errorf("strconv validate query parameters: %w", err)
	}

	if math.IsNaN(f) {
		return fmt.Errorf(queryParam+"%w", errIsNaN)
	}

	return nil
}
