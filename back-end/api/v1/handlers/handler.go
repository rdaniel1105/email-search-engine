package handlers

import (
	"example/mamuro/helpers"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Routes creates a route for searching data.
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", ListEmails)
	// r.Route("/{id}", func(r chi.Router) {
	// 	r.Get("/", GetData())
	// })

	return r
}

// ListEmails displays the list of emails obtained from the request to ZincSearch.
func ListEmails(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("reading request body: %w", err))
		return
	}

	term, err := helpers.ValidateBody(w, requestBody)
	if err != nil {
		fmt.Println(fmt.Errorf("validate body: %w", err))
		return
	}

	query, err := helpers.GetQueryParamsForRequest(r, term)
	if err != nil {
		JSONErrorCheck := helpers.JSONResponse(w, http.StatusOK, map[string]interface{}{"message": err})
		err = helpers.ResponseErrorChecker(JSONErrorCheck, err)

		fmt.Println(fmt.Errorf("Get query for request: %w", err))

		return
	}

	err = helpers.DoRequest(w, query)
	if err != nil {
		fmt.Println(fmt.Errorf("DoRequest: %w", err))
		return
	}
}

// // GetData makes a request to ZincSearch to get the data requested.
// func GetData() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		term := chi.URLParam(r, "id")
// 		requestBody, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		query := helpers.GetQueryForRequest(r, term)

// 		valid := helpers.ValidateBody(requestBody)
// 		if !valid {
// 			fmt.Print("Invalid body from request")
// 			helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"Error": "Request should not have a body."})
// 			return
// 		}

// 		helpers.DoRequest(w, query)
// 	}
// }
