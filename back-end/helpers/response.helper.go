package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RespondWithJSON sends a response in Json format
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
