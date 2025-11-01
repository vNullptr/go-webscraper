package scraper

import (
	"context"
	"fmt"
	"io"       // to read the body
	"net/http" // http requests
	"net/url"  // url handling and type
	"time"
)

// removed httpClient struct its not needed

// need to add custom header handling
// and add timeout handling
func (s *Scraper) FetchURL(rawUrl string, method string, ctx context.Context) (int, error) {

	//creating the request
	if ctx == nil {
		ctx = context.Background()
	}

	// in case i forget we're parsing to make sure its a valid url
	parsedUrl, _ := url.Parse(rawUrl)

	req, err := http.NewRequestWithContext(ctx, method, parsedUrl.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("creating HTTP request failed %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// incase we receive a fail response
		if resp != nil {
			defer resp.Body.Close()
			return resp.StatusCode, fmt.Errorf("failed sending the http request %w", err)
		}

		return 0, fmt.Errorf("failed sending the http request %w", err)
	}

	body, _ := io.ReadAll(resp.Body)
	s.unparsedHTML = body
	return resp.StatusCode, nil

}

func (s *Scraper) FetchURLWithRetry(rawUrl string, method string, timeoutDelay int, limit int) error {
    for i := 0; i < limit; i++ {
        ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutDelay)*time.Second)

        status, err := s.FetchURL(rawUrl, method, ctx)
        cancel() // call per-iteration to avoid leaking resources

        // success -> stop
        if err == nil && status >= 200 && status < 300 {
            return nil
        }

        // non-retryable client error (except 429)
        if err == nil && status >= 400 && status < 500 && status != 429 {
            return fmt.Errorf("non-retryable status %d", status)
        }

        // if not last attempt, wait (simple backoff)
        if i < limit-1 {
            time.Sleep(time.Second * time.Duration(1<<uint(i))) // 1s, 2s, 4s...
        }
    }

    return fmt.Errorf("all %d attempts failed", limit)
}
