package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

const (
	defaultFrom  = 0
	defaultLimit = 25
	maxLimit     = 100

	zincSearchType = "matchphrase"
)

var (
	defaultSourceFields = []string{"From", "To", "Subject", "Body", "Date", "Message-ID"}

	errLimitValueExceeded = fmt.Errorf("the maximum limit is %d", maxLimit)
	errNegativeValue      = errors.New("value must be non-negative")
	errNotANumber         = errors.New("value must be an integer")
)

type searchQuery struct {
	SearchType string     `json:"search_type"`
	Query      searchTerm `json:"query"`
	From       int        `json:"from"`
	MaxResults int        `json:"max_results"`
	Source     []string   `json:"_source"`
}

type searchTerm struct {
	Term string `json:"term"`
}

// GetQueryParamsForRequest parses pagination params from the request and
// builds the ZincSearch query payload.
func GetQueryParamsForRequest(r *http.Request, term string) (string, error) {
	from, err := parsePaginationParam(r.URL.Query().Get("from"), defaultFrom)
	if err != nil {
		return "", fmt.Errorf("from: %w", err)
	}

	limit, err := parsePaginationParam(r.URL.Query().Get("limit"), defaultLimit)
	if err != nil {
		return "", fmt.Errorf("limit: %w", err)
	}

	if limit > maxLimit {
		return "", errLimitValueExceeded
	}

	return SetQueryForRequest(from, limit, term)
}

// SetQueryForRequest builds the ZincSearch search request body by marshaling a
// typed struct, which prevents injection through the term parameter.
func SetQueryForRequest(from, limit int, term string) (string, error) {
	q := searchQuery{
		SearchType: zincSearchType,
		Query:      searchTerm{Term: term},
		From:       from,
		MaxResults: limit,
		Source:     defaultSourceFields,
	}

	body, err := json.Marshal(q)
	if err != nil {
		return "", fmt.Errorf("marshal search query: %w", err)
	}

	return string(body), nil
}

func parsePaginationParam(raw string, defaultValue int) (int, error) {
	if raw == "" {
		return defaultValue, nil
	}

	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, errNotANumber
	}

	if n < 0 {
		return 0, errNegativeValue
	}

	return n, nil
}
