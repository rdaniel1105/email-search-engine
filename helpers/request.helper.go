package helpers

import (
	"encoding/json"
	"example/mamuro/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type reqHeaders struct {
	admin       string
	password    string
	ContentType string
	UserAgent   string
}

var requestHeaders = reqHeaders{
	"admin",
	"Complexpass#123",
	"application/json",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36",
}

// EmailResponse copies the underlying structure coming from models
type EmailResponse models.EmailResponse

// DoRequest performs a request to zincsearch
func DoRequest(query string, w http.ResponseWriter) {
	var email *EmailResponse

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	zincSearchURL := os.Getenv("ZincSearchURL")

	req, err := http.NewRequest(http.MethodPost, zincSearchURL, strings.NewReader(query))
	if err != nil {
		fmt.Print(err)
	}

	req.SetBasicAuth(requestHeaders.admin, requestHeaders.password)
	req.Header.Set("Content-Type", requestHeaders.ContentType)
	req.Header.Set("User-Agent", requestHeaders.UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"InternalServerError": "Could not connect to database."})

		return
	}

	email, err = ZincSearchResponseStatus(resp)
	if err != nil {
		RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("%s", err)})

		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{"total": email.Hits.Total, "hits": email.Hits.Hits})

}

// ZincSearchResponseStatus ensures we're getting the proper response status from the database.
func ZincSearchResponseStatus(httpResponse *http.Response) (*EmailResponse, error) {
	statusResponse := &EmailResponse{HTTPResponse: httpResponse}

	defer closeResponseBody(httpResponse)

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println("Error reading from API:", err)
		return statusResponse, err
	}

	err = json.Unmarshal(body, statusResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return statusResponse, err
	}

	return statusResponse, nil
}

func closeResponseBody(response *http.Response) {
	err := response.Body.Close()
	if err != nil {
		fmt.Println("Error closing response body:", err)
	}
}
