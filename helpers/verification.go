package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TokenVerification encapsulates the response from Authy API when verifying a token.
type TokenVerification struct {
	HTTPResponse *http.Response
	Message      string      `json:"message"`
	Token        string      `json:"token"`
	Success      interface{} `json:"success"`
}

// NewTokenVerification creates an instance of a TokenVerification
func NewTokenVerification(response *http.Response) (*TokenVerification, error) {
	tokenVerification := &TokenVerification{HTTPResponse: response}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading from API:", err)
		return tokenVerification, err
	}

	err = json.Unmarshal(body, &tokenVerification)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return tokenVerification, err
	}

	return tokenVerification, nil
}

// Valid returns true if the verification was valid.
func (verification *TokenVerification) Valid() bool {
	if verification.HTTPResponse.StatusCode == 200 && verification.Token == "is valid" {
		return true
	}

	return false
}
