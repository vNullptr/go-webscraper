package scraper

import (
	"context"
	"fmt"
	"io"       // to read the body
	"net/http" // http requests
	"net/url"  // url handling and type
)

// removed httpClient struct its not needed

// need to use http.NewRequest for other methods and custom headers
// bonus : add timeout handling
func (s *Scraper) DlUrl(rawUrl string, method string) ([]byte, error) {

	//creating the request
	ctx := context.Background()

	// in case i forget we're parsing to make sure its a valid url
	parsedUrl, _ := url.Parse(rawUrl)

	req, err := http.NewRequestWithContext(ctx, method, parsedUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("creating HTTP request failed %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed sending the http request %w", err)
	}

	body, _ := io.ReadAll(resp.Body)
	return body, nil

}
