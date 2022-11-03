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

	r.Get("/", List)
	// r.Route("/{input}", func(r chi.Router) {
	// 	r.Get("/", GetData())
	// })

	return r
}

// List displays the list of emails obtained from the request to ZincSearch.
func List(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	term, valid := helpers.ValidateBody(w, requestBody)
	if !valid {
		fmt.Print("Invalid body from request")
		return
	}
	query := helpers.GetQueryForRequest(r, term)
	// query := `{
	//     "search_type": "alldocuments",
	//     "query":
	//     {"term": ""},
	//     "_source": ["From","To","Subject", "Body","Date","Message-ID"]
	// }`

	helpers.DoRequest(w, query)
}

// // GetData makes a request to ZincSearch to get the data requested.
// func GetData() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		term := chi.URLParam(r, "input")
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
