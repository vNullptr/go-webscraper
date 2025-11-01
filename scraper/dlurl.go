package scraper

import (
	"context"
	"fmt"
	"io"       // to read the body
	"net/http" // http requests
	"net/url"  // url handling and type
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
