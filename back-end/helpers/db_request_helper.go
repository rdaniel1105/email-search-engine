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

var (
	errDBRequest  = errors.New("internal server error, could not connect to database")
	errDBResponse = errors.New("internal server error, unexpected response from database")
)

// EmailResponse copies the underlying structure coming from models
type EmailResponse models.EmailResponse

// DoRequest performs a request to zincsearch
func DoRequest(w http.ResponseWriter, query string) error {
	var email *EmailResponse

	zincSearchURL := os.Getenv("ZincSearchURL")

	req, err := http.NewRequest(http.MethodPost, zincSearchURL, strings.NewReader(query))
	if err != nil {
		return fmt.Errorf("newrequest wrapping: %w", err)
	}

	req.SetBasicAuth(requestHeaders.admin, requestHeaders.password)
	req.Header.Set("Content-Type", requestHeaders.ContentType)
	req.Header.Set("User-Agent", requestHeaders.UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ResponseErrorHelper(w, http.StatusInternalServerError, errDBRequest.Error(), fmt.Errorf("zincSearch request: %w", err))
	}

	email, err = ZincSearchResponseStatus(resp)
	if err != nil {
		return ResponseErrorHelper(w, http.StatusInternalServerError, errDBResponse.Error(), fmt.Errorf("zincSearch response: %w", err))
	}

	JSONErrorCheck := JSONResponse(w, http.StatusOK, map[string]interface{}{"total": email.Hits.Total, "hits": email.Hits.Hits})
	err = ResponseErrorChecker(JSONErrorCheck, nil)

	return err
}

// ZincSearchResponseStatus ensures we're getting the proper response status from the database.
func ZincSearchResponseStatus(httpResponse *http.Response) (*EmailResponse, error) {
	statusResponse := &EmailResponse{HTTPResponse: httpResponse}

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
