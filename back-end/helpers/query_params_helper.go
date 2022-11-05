package helpers

import "net/http"

// GetQueryParamsForRequest returns all relevant query parameters.
func GetQueryParamsForRequest(r *http.Request, term string) (string, error) {
	var query string

	if from, limit := r.URL.Query().Get("from"), r.URL.Query().Get("limit"); from != "" && limit != "" {
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

	from := r.URL.Query().Get("from")
	if from != "" {
		err := ValidateQueryParams(from)
		if err != nil {
			return "", err
		}
	} else {
		from = "0"
	}

	limit := r.URL.Query().Get("limit")
	if limit != "" {
		err := ValidateQueryParams(limit)
		if err != nil {
			return "", err
		}
	} else {
		limit = "25"
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
