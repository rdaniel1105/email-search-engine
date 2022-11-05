package models

import "net/http"

// TODO mejorar structs

// EmailSource represents the source contained in the Email struct
type EmailSource struct {
	Body      string `json:"Body"`
	Date      string `json:"Date"`
	From      string `json:"From"`
	MessageID string `json:"Message-ID"`
	Subject   string `json:"Subject"`
	To        string `json:"To"`
}

//Email represents an email struct
type Email struct {
	Index     string      `json:"_index"`
	Type      string      `json:"_type"`
	ID        string      `json:"_id"`
	Score     float64     `json:"_score"`
	TimeStamp string      `json:"@timestamp"`
	Source    EmailSource `json:"_source"`
}

// EmailResponseShards represents shards in EmailResponse
type EmailResponseShards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

// EmailResponseHitsTotal represents Total in EmailResponse's Hits
type EmailResponseHitsTotal struct {
	Value int `json:"value"`
}

// EmailResponseHits represents Hits in EmailResponse
type EmailResponseHits struct {
	Total    EmailResponseHitsTotal `json:"total"`
	MaxScore float64                `json:"max_score"`
	Hits     []Email                `json:"hits"`
}

// EmailResponse represents the expected response from zincsearch
type EmailResponse struct {
	HTTPResponse *http.Response
	Took         int                 `json:"took"`
	Timedoout    bool                `json:"timed_out"`
	Shards       EmailResponseShards `json:"_shards"`
	Hits         EmailResponseHits   `json:"hits"`
}
