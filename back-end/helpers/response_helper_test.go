package helpers

import (
	"math"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestJSONResponse does asdasdas
func TestJSONResponse(t *testing.T) {
	err := JSONResponse(nil, http.StatusAccepted, math.NaN())
	assert.Contains(t, err.Error(), "marshal json response")
}
