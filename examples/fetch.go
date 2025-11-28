package main

import (
	"fmt"

	"webscraper/pkg/scraper"
)

func main() {

	s := scraper.CreateScraper();
	buffer, _, err := scraper.FetchURL("https://example.com", "GET", nil)
	if (err == nil && len(buffer) > 0) {
		s.SetUnparsedHTML(buffer)
	}

	fmt.Println(s.GetUnparsedHTML())
}
