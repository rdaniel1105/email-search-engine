package helpers

import (
	"math"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONResponse(t *testing.T) {
	err := JSONResponse(nil, http.StatusAccepted, math.NaN())
	assert.Contains(t, err.Error(), "json marshal response")
}
