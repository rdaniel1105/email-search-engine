package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// BodyRegex is the supported body
var (
	BodyRegex  = regexp.MustCompile(`^[a-zA-Z]*$`)
	bodyParsed struct {
		Term string `json:"term"`
	}
)

var (
	errMissingBody    = errors.New("missing body")
	errMissingTermKey = errors.New(`"term" key missing`)
	errInvalidBody    = errors.New("invalid body")
)

// ValidateBody makes sure the body field is empty
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

	if !BodyRegex.MatchString(bodyParsed.Term) {
		return "", ResponseErrorHelper(w, http.StatusBadRequest, errInvalidBody.Error(), errInvalidBody)
	}

	return bodyParsed.Term, nil
}

// ValidateQueryParams validates that the query params are numbers
func ValidateQueryParams(queryParam string) error {
	f, err := strconv.ParseFloat(queryParam, 64)
	if err != nil {
		return fmt.Errorf("strconv validate query parameters: %w", err)
	}

	if math.IsNaN(f) {
		return errors.New(queryParam + `is not a number`)
	}
	return nil
}
