package models

import "net/http"

//Email represents an email struct
type Email struct {
	Index     string  `json:"_index"`
	Type      string  `json:"_type"`
	ID        string  `json:"_id"`
	Score     float64 `json:"_score"`
	TimeStamp string  `json:"@timestamp"`
	Source    struct {
		Body      string `json:"Body"`
		Date      string `json:"Date"`
		From      string `json:"From"`
		MessageID string `json:"Message-ID"`
		Subject   string `json:"Subject"`
		To        string `json:"To"`
	} `json:"_source"`
}

// EmailResponse represents the expected response from zincsearch
type EmailResponse struct {
	HTTPResponse *http.Response
	Took         int  `json:"took"`
	Timedoout    bool `json:"timed_out"`
	Shards       struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []Email `json:"hits"`
	} `json:"hits"`
}
