package models

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToSearchResponse_Empty(t *testing.T) {
	c := require.New(t)

	got := ToSearchResponse(nil)
	c.Equal(0, got.Total)
	c.NotNil(got.Hits)
	c.Len(got.Hits, 0)
}

func TestToSearchResponse_MapsFlatly(t *testing.T) {
	c := require.New(t)

	raw := &EmailResponse{
		Hits: EmailResponseHits{
			Total: EmailResponseHitsTotal{Value: 2},
			Hits: []Email{
				{
					ID:    "<a@thyme>",
					Score: 1.5,
					Source: EmailSource{
						From:    "alice@example.com",
						To:      "bob@example.com",
						Subject: "hello",
						Date:    "Wed, 10 Oct 2001",
						Body:    "world",
					},
				},
				{
					ID:    "<b@thyme>",
					Score: 2.5,
					Source: EmailSource{
						From:    "carol@example.com",
						Subject: "ping",
					},
				},
			},
		},
	}

	got := ToSearchResponse(raw)
	c.Equal(2, got.Total)
	c.Len(got.Hits, 2)

	c.Equal("<a@thyme>", got.Hits[0].ID)
	c.Equal(1.5, got.Hits[0].Score)
	c.Equal("alice@example.com", got.Hits[0].From)
	c.Equal("bob@example.com", got.Hits[0].To)
	c.Equal("hello", got.Hits[0].Subject)
	c.Equal("world", got.Hits[0].Body)

	c.Equal("<b@thyme>", got.Hits[1].ID)
	c.Equal("carol@example.com", got.Hits[1].From)
	c.Empty(got.Hits[1].To)
}

// Guards against the ZincSearch internals (_index, _type, @timestamp,
// _source.*) reappearing in the public API response.
func TestSearchResponse_JSON_DoesNotLeakStorageFields(t *testing.T) {
	c := require.New(t)

	resp := SearchResponse{
		Total: 1,
		Hits: []SearchHit{
			{ID: "<x@thyme>", From: "alice@example.com", Subject: "hi"},
		},
	}

	body, err := json.Marshal(resp)
	c.NoError(err)

	encoded := string(body)
	for _, leaked := range []string{`"_index"`, `"_type"`, `"_id"`, `"_score"`, `"@timestamp"`, `"_source"`, `"Message-ID"`} {
		c.False(strings.Contains(encoded, leaked), "encoded response unexpectedly contains %q: %s", leaked, encoded)
	}

	// Public keys are present and snake-cased.
	for _, want := range []string{`"id"`, `"from"`, `"subject"`, `"total"`, `"hits"`} {
		c.True(strings.Contains(encoded, want), "encoded response missing key %q: %s", want, encoded)
	}
}
