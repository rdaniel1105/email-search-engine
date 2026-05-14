package models

import "net/http"

// ---------------------------------------------------------------------------
// Public API shape — what the back-end returns to the client.
// ---------------------------------------------------------------------------

// SearchHit is the per-result row returned to the client. Flat, snake-case,
// no leaking of ZincSearch's storage fields (_index, _type, @timestamp, etc.).
type SearchHit struct {
	ID      string  `json:"id"`
	Score   float64 `json:"score"`
	From    string  `json:"from"`
	To      string  `json:"to"`
	Subject string  `json:"subject"`
	Date    string  `json:"date"`
	Body    string  `json:"body"`
}

// SearchResponse is the top-level shape the API returns for a search.
type SearchResponse struct {
	Total int         `json:"total"`
	Hits  []SearchHit `json:"hits"`
}

// ToSearchResponse maps the raw ZincSearch response into the public shape.
func ToSearchResponse(raw *EmailResponse) SearchResponse {
	if raw == nil {
		return SearchResponse{Hits: []SearchHit{}}
	}

	out := SearchResponse{
		Total: raw.Hits.Total.Value,
		Hits:  make([]SearchHit, 0, len(raw.Hits.Hits)),
	}

	for _, h := range raw.Hits.Hits {
		out.Hits = append(out.Hits, SearchHit{
			ID:      h.ID,
			Score:   h.Score,
			From:    h.Source.From,
			To:      h.Source.To,
			Subject: h.Source.Subject,
			Date:    h.Source.Date,
			Body:    h.Source.Body,
		})
	}

	return out
}

// ---------------------------------------------------------------------------
// Internal ZincSearch wire format — used by the helpers layer to decode the
// upstream response. Not exported through the API.
// ---------------------------------------------------------------------------

// EmailSource represents the source contained in the Email struct.
type EmailSource struct {
	Body      string `json:"Body"`
	Date      string `json:"Date"`
	From      string `json:"From"`
	MessageID string `json:"Message-ID"`
	Subject   string `json:"Subject"`
	To        string `json:"To"`
}

// Email represents an email struct as returned by ZincSearch.
type Email struct {
	Index     string      `json:"_index"`
	Type      string      `json:"_type"`
	ID        string      `json:"_id"`
	Score     float64     `json:"_score"`
	TimeStamp string      `json:"@timestamp"`
	Source    EmailSource `json:"_source"`
}

// EmailResponseShards represents shards in EmailResponse.
type EmailResponseShards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

// EmailResponseHitsTotal represents Total in EmailResponse's Hits.
type EmailResponseHitsTotal struct {
	Value int `json:"value"`
}

// EmailResponseHits represents Hits in EmailResponse.
type EmailResponseHits struct {
	Total    EmailResponseHitsTotal `json:"total"`
	MaxScore float64                `json:"max_score"`
	Hits     []Email                `json:"hits"`
}

// EmailResponse represents the expected response from ZincSearch.
type EmailResponse struct {
	HTTPResponse *http.Response
	Took         int                 `json:"took"`
	Timedoout    bool                `json:"timed_out"`
	Shards       EmailResponseShards `json:"_shards"`
	Hits         EmailResponseHits   `json:"hits"`
}
