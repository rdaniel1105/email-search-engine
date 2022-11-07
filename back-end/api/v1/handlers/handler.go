package handlers

import (
	"errors"
	"example/mamuro/helpers"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var (
	errReadingRequestBody     = errors.New("reading request body")
	errValidatingBody         = errors.New("validate body")
	errGettingQueryForRequest = errors.New("get query for request")
	errDoingRequest           = errors.New("DoRequest")
)

// Routes creates a route for searching data in the API.
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", ListEmails)

	return r
}

// ListEmails displays the list of emails obtained from the request to ZincSearch.
func ListEmails(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(errReadingRequestBody, err)
		return
	}

	term, err := helpers.ValidateBody(w, requestBody)
	if err != nil {
		fmt.Println(errValidatingBody, err)
		return
	}

	query, err := helpers.GetQueryParamsForRequest(r, term)
	if err != nil {
		JSONErrorCheck := helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{"message": err})
		err = helpers.ResponseErrorChecker(JSONErrorCheck, err)

		fmt.Println(errGettingQueryForRequest, err)

		return
	}

	err = helpers.DoRequest(w, query)
	if err != nil {
		fmt.Println(errDoingRequest, err)
		return
	}
}
