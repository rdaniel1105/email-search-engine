package helpers

import (
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetQueryForRequest_EscapesInjection(t *testing.T) {
	c := require.New(t)

	// A term containing a double quote and JSON-control characters must be
	// safely escaped by json.Marshal, not break out of the "term" field.
	maliciousTerm := `"},"max_results":99999,"extra":"`
	body, err := SetQueryForRequest(0, 10, maliciousTerm)
	c.NoError(err)

	var round map[string]interface{}
	c.NoError(json.Unmarshal([]byte(body), &round))
	c.Equal(float64(10), round["max_results"]) // attacker's value was not honored

	query := round["query"].(map[string]interface{})
	c.Equal(maliciousTerm, query["term"])
}

func TestParsePaginationParam(t *testing.T) {
	c := require.New(t)

	n, err := parsePaginationParam("", 25)
	c.NoError(err)
	c.Equal(25, n)

	n, err = parsePaginationParam("7", 25)
	c.NoError(err)
	c.Equal(7, n)

	_, err = parsePaginationParam("-1", 25)
	c.ErrorIs(err, errNegativeValue)

	_, err = parsePaginationParam("abc", 25)
	c.ErrorIs(err, errNotANumber)

	_, err = parsePaginationParam("1.5", 25)
	c.ErrorIs(err, errNotANumber)
}

func TestGetQueryParamsForRequest_AppliesDefaults(t *testing.T) {
	c := require.New(t)

	r := httptest.NewRequest("POST", "/search", nil)
	body, err := GetQueryParamsForRequest(r, "abc")
	c.NoError(err)

	var q map[string]interface{}
	c.NoError(json.Unmarshal([]byte(body), &q))
	c.Equal(float64(defaultFrom), q["from"])
	c.Equal(float64(defaultLimit), q["max_results"])
}

func TestGetQueryParamsForRequest_EnforcesMaxLimit(t *testing.T) {
	c := require.New(t)

	r := httptest.NewRequest("POST", "/search?from=0&limit="+strconv.Itoa(maxLimit+1), nil)
	_, err := GetQueryParamsForRequest(r, "abc")
	c.ErrorIs(err, errLimitValueExceeded)
}

func TestGetQueryParamsForRequest_RejectsNegative(t *testing.T) {
	c := require.New(t)

	r := httptest.NewRequest("POST", "/search?from=-5&limit=10", nil)
	_, err := GetQueryParamsForRequest(r, "abc")
	c.ErrorIs(err, errNegativeValue)
}
