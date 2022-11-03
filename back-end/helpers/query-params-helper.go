package helpers

import "net/http"

// GetQueryForRequest returns all relevant query parameters.
func GetQueryForRequest(r *http.Request, term string) string {
	from := r.URL.Query().Get("from")
	limit := r.URL.Query().Get("limit")
	query := SetQueryForRequest(from, limit, term)

	return query
}

// SetQueryForRequest sets the correct query for the request to the database
func SetQueryForRequest(from string, limit string, term string) string {
	var query string

	switch {
	case from != "" && limit != "":
		query = `{
			"search_type": "matchphrase",
			"query":
			{
			"term": "` + term + `"
			},
			"from": ` + from + `,
        	"max_results": ` + limit + `,
			"_source": ["From","To","Subject", "Body","Date","Message-ID"]
			}`
	case limit != "" && from == "":
		query = `{
			"search_type": "matchphrase",
			"query":
			{
			"term": "` + term + `"
			},
        	"max_results": ` + limit + `,
			"_source": ["From","To","Subject", "Body","Date","Message-ID"]
			}`
	case limit == "" && from != "":
		query = `{
			"search_type": "matchphrase",
			"query":
			{
			"term": "` + term + `"
			},
			"from": ` + from + `,
			"_source": ["From","To","Subject", "Body","Date","Message-ID"]
			}`
	default:
		query = `{
			"search_type": "matchphrase",
			"query":
			{
			"term": "` + term + `"
			},
			"_source": ["From","To","Subject", "Body","Date","Message-ID"]
			}`
	}
	return query
}
