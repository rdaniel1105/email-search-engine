package helpers

import (
	"math"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONResponse(t *testing.T) {
	c := require.New(t)

	err := JSONResponse(nil, http.StatusAccepted, math.NaN())
	c.ErrorContains(err, "json marshal response")
}
