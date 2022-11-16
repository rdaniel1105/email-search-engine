package helpers

import (
	"encoding/json"
	"example/mamuro/models"
	"fmt"
	"net/http"
)

// JSONResponse sends a response in Json format
func JSONResponse(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("json marshal response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("json writer response: %w", err)
	}

	return nil
}

// ResponseErrorChecker verifies if an error ocurred with JSONResponse, if so, returns that error, otherwise it returns the originar error.
func ResponseErrorChecker(JSONResponseErr error, err error) error {
	if JSONResponseErr == nil {
		return err
	}

	return JSONResponseErr
}

// ResponseErrorHelper calls and checks the JSONResponse func an returns the corresponding error
func ResponseErrorHelper(w http.ResponseWriter, code int, err error) error {
	JSONErrorCheck := JSONResponse(w, code, models.ErrorResponseMessage{"message": err.Error()})

	return ResponseErrorChecker(JSONErrorCheck, err)
}
