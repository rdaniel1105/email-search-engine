package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateBody_EmptyBody(t *testing.T) {
	c := require.New(t)

	term, err := ValidateBody(nil)
	c.ErrorIs(err, errMissingBody)
	c.Empty(term)
}

func TestValidateBody_InvalidJSON(t *testing.T) {
	c := require.New(t)

	term, err := ValidateBody([]byte("{not json"))
	c.ErrorIs(err, errInvalidJSON)
	c.Empty(term)
}

func TestValidateBody_MissingTerm(t *testing.T) {
	c := require.New(t)

	term, err := ValidateBody([]byte(`{"other":"value"}`))
	c.ErrorIs(err, errMissingTermKey)
	c.Empty(term)
}

func TestValidateBody_EmptyTerm(t *testing.T) {
	c := require.New(t)

	term, err := ValidateBody([]byte(`{"term":""}`))
	c.ErrorIs(err, errMissingTermKey)
	c.Empty(term)
}

func TestValidateBody_Valid(t *testing.T) {
	c := require.New(t)

	term, err := ValidateBody([]byte(`{"term":"hello world"}`))
	c.NoError(err)
	c.Equal("hello world", term)
}

func TestValidateBody_IgnoresExtraFields(t *testing.T) {
	c := require.New(t)

	term, err := ValidateBody([]byte(`{"term":"abc","ignored":42}`))
	c.NoError(err)
	c.Equal("abc", term)
}

// Injection attempt: the original implementation concatenated the term into
// the JSON query string, so a quote-containing term broke out. ValidateBody
// must return it as-is and the marshaler in query_params.go must safely
// escape it. This test guards the first half.
func TestValidateBody_PreservesQuotesInTerm(t *testing.T) {
	c := require.New(t)

	term, err := ValidateBody([]byte(`{"term":"quote\" injection"}`))
	c.NoError(err)
	c.Equal(`quote" injection`, term)
}
