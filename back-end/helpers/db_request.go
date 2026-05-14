package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rdaniel1105/email-search-engine/back-end/models"
)

const (
	defaultZincSearchURL   = "http://localhost:4080"
	defaultZincSearchIndex = "emails"

	contentTypeJSON = "application/json"
	userAgent       = "email-search-engine"

	dbRequestTimeout = 10 * time.Second
)

var (
	httpClient = &http.Client{Timeout: dbRequestTimeout}

	errDBUnreachable = errors.New("search database unavailable")
	errDBResponse    = errors.New("unexpected response from search database")
)

// DoRequest performs a search against ZincSearch and returns the parsed
// response. Internal errors (connection, decoding, upstream status) are
// logged and surfaced as generic sentinel errors so callers can map them
// to a safe client-facing message without leaking infrastructure details.
func DoRequest(query string) (*models.EmailResponse, error) {
	baseURL := envOr("ZINCSEARCH_URL", defaultZincSearchURL)
	index := envOr("ZINCSEARCH_INDEX", defaultZincSearchIndex)

	dbURL := fmt.Sprintf("%s/api/%s/_search", baseURL, index)

	req, err := http.NewRequest(http.MethodPost, dbURL, strings.NewReader(query))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	req.SetBasicAuth(os.Getenv("ZINCSEARCH_USERNAME"), os.Getenv("ZINCSEARCH_PASSWORD"))
	req.Header.Set("Content-Type", contentTypeJSON)
	req.Header.Set("User-Agent", userAgent)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("zincsearch request: %v", err)
		return nil, errDBUnreachable
	}
	defer closeResponseBody(resp)

	parsed, err := parseSearchResponse(resp)
	if err != nil {
		log.Printf("zincsearch response: %v", err)
		return nil, errDBResponse
	}

	return parsed, nil
}

func parseSearchResponse(resp *http.Response) (*models.EmailResponse, error) {
	out := &models.EmailResponse{HTTPResponse: resp}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	if err := json.Unmarshal(body, out); err != nil {
		return nil, fmt.Errorf("decode body: %w", err)
	}

	return out, nil
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

func closeResponseBody(resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		log.Printf("close response body: %v", err)
	}
}
