package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	errLimitValueExceeded = errors.New("the maximux limit is 100")
)

// GetQueryParamsForRequest returns all relevant query parameters.
func GetQueryParamsForRequest(r *http.Request, term string) (string, error) {
	var query string

	from, limit := r.URL.Query().Get("from"), r.URL.Query().Get("limit")
	if from != "" && limit != "" {
		err := ValidateQueryParams(from)
		if err != nil {
			return "", err
		}

		err = ValidateQueryParams(limit)
		if err != nil {
			return "", err
		}

		query = SetQueryForRequest(from, limit, term)

		return query, nil
	}

	if from == "" {
		from = "0"
	}

	if limit == "" {
		limit = "25"
	}

	err := ValidateQueryParams(from)
	if err != nil {
		return "", err
	}

	err = ValidateQueryParams(limit)
	if err != nil {
		return "", err
	}

	limit, err = limitChecker(limit)
	if err != nil {
		return "", err
	}

	query = SetQueryForRequest(from, limit, term)

	return query, nil
}

// SetQueryForRequest sets the correct query for the request to the database
func SetQueryForRequest(from string, limit string, term string) string {

	query := `{
			"search_type": "matchphrase",
			"query":
			{
			"term": "` + term + `"
			},
			"from": ` + from + `,
        	"max_results": ` + limit + `,
			"_source": ["From","To","Subject", "Body","Date","Message-ID"]
			}`

	return query
}

func limitChecker(limit string) (string, error) {
	numLimit, err := strconv.ParseInt(limit, 36, 64)
	if err != nil {
		return "", fmt.Errorf("limitChecker: %w", err)
	}

	if numLimit > 100 {
		return "", errLimitValueExceeded
	}

	return strconv.FormatInt(numLimit, 36), nil
}
