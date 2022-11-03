package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// BodyRegex is the supported body
var (
	BodyRegex  = regexp.MustCompile(`^[a-zA-Z]*$`)
	BodyParsed struct {
		Term string `json:"term"`
	}
)

// ValidateBody makes sure the body field is empty
func ValidateBody(w http.ResponseWriter, body []byte) (string, bool) {
	err := json.Unmarshal(body, &BodyParsed)
	if err != nil {
		fmt.Println(err)
	}

	if string(body) == "" || !strings.Contains(string(body), "term") {
		RespondWithJSON(w, http.StatusBadRequest, map[string]string{"Message": `Empty body or key is different than "term"`})
		return BodyParsed.Term, false
	}

	if !BodyRegex.MatchString(BodyParsed.Term) {
		return BodyParsed.Term, false
	}

	return BodyParsed.Term, true
}

// ValidateID validates the id provided
func ValidateID() bool {
	return true
}
