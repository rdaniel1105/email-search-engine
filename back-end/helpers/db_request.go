package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rdaniel1105/email-search-engine/back-end/models"
)

type reqHeaders struct {
	contentType string
	userAgent   string
}

var (
	requestHeaders = reqHeaders{
		contentType: "application/json",
		userAgent:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36",
	}

	defaultZincSearchURL   = "http://localhost:4080"
	defaultZincSearchIndex = "emails"
)

const (
	errDBRequest  = "DB request, could not connect to database: %w"
	errDBResponse = "DB response, unexpected response from database: %w"
)

// DoRequest performs a request to the DB
func DoRequest(w http.ResponseWriter, query string) error {
	var matchedEmails *models.EmailResponse

	baseURL := os.Getenv("ZINCSEARCH_URL")
	if baseURL == "" {
		baseURL = defaultZincSearchURL
	}

	index := os.Getenv("ZINCSEARCH_INDEX")
	if index == "" {
		index = defaultZincSearchIndex
	}

	dbURL := fmt.Sprintf("%s/api/%s/_search", baseURL, index)

	admin := os.Getenv("ZINCSEARCH_USERNAME")
	password := os.Getenv("ZINCSEARCH_PASSWORD")

	req, err := http.NewRequest(http.MethodPost, dbURL, strings.NewReader(query))
	if err != nil {
		return fmt.Errorf("newrequest wrapping: %w", err)
	}

	req.SetBasicAuth(admin, password)
	req.Header.Set("Content-Type", requestHeaders.contentType)
	req.Header.Set("User-Agent", requestHeaders.userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ResponseErrorHelper(w, http.StatusInternalServerError, fmt.Errorf(errDBRequest, err))
	}

	matchedEmails, err = DataBaseResponseStatus(resp)
	if err != nil {
		return ResponseErrorHelper(w, http.StatusInternalServerError, fmt.Errorf(errDBResponse, err))
	}

	JSONErrorCheck :=
		JSONResponse(w, http.StatusOK, map[string]interface{}{"DBresponse": matchedEmails.HTTPResponse.StatusCode, "total": matchedEmails.Hits.Total, "hits": matchedEmails.Hits.Hits})

	return ResponseErrorChecker(JSONErrorCheck, nil)
}

// DataBaseResponseStatus checks if we're getting the proper response status from the database.
func DataBaseResponseStatus(httpResponse *http.Response) (*models.EmailResponse, error) {
	statusResponse := &models.EmailResponse{HTTPResponse: httpResponse}

	defer closeResponseBody(httpResponse)

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println("reading from API:", err)
		return statusResponse, err
	}

	err = json.Unmarshal(body, statusResponse)
	if err != nil {
		fmt.Println("parsing JSON:", err)
		return statusResponse, err
	}

	return statusResponse, nil
}

func closeResponseBody(response *http.Response) {
	err := response.Body.Close()
	if err != nil {
		fmt.Println("error closing response body:", err)
	}
}
