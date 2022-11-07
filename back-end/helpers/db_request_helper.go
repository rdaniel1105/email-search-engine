package helpers

import (
	"encoding/json"
	"errors"
	"example/mamuro/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

	zincSearchURLSet = "http://localhost:4080/api/myindex/_search"

	errDBRequest  = errors.New("could not connect to database")
	errDBResponse = errors.New("unexpected response from database")
)

// DoRequest performs a request to zincsearch
func DoRequest(w http.ResponseWriter, query string) error {
	var matchedEmails *models.EmailResponse

	zincSearchURL := os.Getenv("ZincSearchURL")
	if zincSearchURL == "" {
		zincSearchURL = zincSearchURLSet
	}

	admin := os.Getenv("ADMIN")
	password := os.Getenv("PASSWORD")

	req, err := http.NewRequest(http.MethodPost, zincSearchURL, strings.NewReader(query))
	if err != nil {
		return fmt.Errorf("newrequest wrapping: %w", err)
	}

	req.SetBasicAuth(admin, password)
	req.Header.Set("Content-Type", requestHeaders.contentType)
	req.Header.Set("User-Agent", requestHeaders.userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ResponseErrorHelper(w, http.StatusInternalServerError, errDBRequest.Error(), fmt.Errorf("zincSearch request: %w", err))
	}

	matchedEmails, err = DataBaseResponseStatus(resp)
	if err != nil {
		return ResponseErrorHelper(w, http.StatusInternalServerError, errDBResponse.Error(), fmt.Errorf("zincSearch response: %w", err))
	}

	JSONErrorCheck :=
		JSONResponse(w, http.StatusOK, map[string]interface{}{"total": matchedEmails.Hits.Total, "hits": matchedEmails.Hits.Hits})
	err = ResponseErrorChecker(JSONErrorCheck, nil)

	return err
}

// DataBaseResponseStatus ensures we're getting the proper response status from the database.
func DataBaseResponseStatus(httpResponse *http.Response) (*models.EmailResponse, error) {
	statusResponse := &models.EmailResponse{HTTPResponse: httpResponse}

	defer closeResponseBody(httpResponse)

	body, err := ioutil.ReadAll(httpResponse.Body)
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
