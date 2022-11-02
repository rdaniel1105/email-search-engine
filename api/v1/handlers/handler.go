package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//Email represents an email struct
type Email struct {
	Index     string  `json:"_index"`
	Type      string  `json:"_type"`
	ID        string  `json:"_id"`
	Score     float64 `json:"_score"`
	TimeStamp string  `json:"@timestamp"`
	Source    struct {
		Body      string `json:"Body"`
		Date      string `json:"Date"`
		From      string `json:"From"`
		MessageID string `json:"Message-ID"`
		Subject   string `json:"Subject"`
		To        string `json:"To"`
	} `json:"_source"`
}

// EmailResponse represents the expected response from zincsearch
type EmailResponse struct {
	Took      int  `json:"took"`
	Timedoout bool `json:"timed_out"`
	Shards    struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []Email `json:"hits"`
	} `json:"hits"`
}

// Routes creates a route for searching data.
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", List)
	r.Route("/{input}", func(r chi.Router) {
		r.With(ValidateURL).Get("/", GetData())
	})

	return r
}

// ValidateURL validates the URL for the parameter route
func ValidateURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		inputParam := chi.URLParam(r, "input")
		if inputParam == "" {
			w.WriteHeader(404)
			_, err := w.Write([]byte(http.StatusText(404)))
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

// List provides a displays the list of emails obtained from the request to ZincSearch.
func List(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var email EmailResponse
	zincSearchURL := os.Getenv("ZincSearchURL")

	query := `{
        "search_type": "alldocuments",
        "query":
        {"term": ""},
        "_source": ["From","To","Subject", "Body","Date","Message-ID"]
    }`

	req, err := http.NewRequest(http.MethodPost, zincSearchURL, strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer closeResponseBody(resp)

	err = json.Unmarshal(body, &email)
	if err != nil {
		log.Fatal(err)
	}

	respondwithJSON(w, http.StatusOK, email.Hits.Hits)
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	//fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}

// GetData makes a request to ZincSearch to get the data requested.
func GetData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		var email EmailResponse
		zincSearchURL := os.Getenv("ZincSearchURL")
		term := chi.URLParam(r, "input")

		//validate term with Regex
		query := `{
        "search_type": "match",
        "query":
        {
            "term": "` + term + `"
        },
        "_source": ["From","To","Subject", "Body","Date","Message-ID"]
    	}`

		fmt.Printf("%+v\n", query)

		req, err := http.NewRequest(http.MethodPost, zincSearchURL, strings.NewReader(query))
		if err != nil {
			log.Fatal(err)
		}

		req.SetBasicAuth("admin", "Complexpass#123")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer closeResponseBody(resp)

		log.Println(resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(body, &email)
		if err != nil {
			log.Fatal(err)
		}

		respondwithJSON(w, http.StatusOK, email.Hits.Hits)
	}
}

func closeResponseBody(response *http.Response) {
	err := response.Body.Close()
	if err != nil {
		fmt.Println("Error closing response body:", err)
	}
}
