package main

import (
	"fmt"

	"webscraper/pkg/scraper"
)

func main() {

	s := scraper.CreateScraper();
	s.FetchURL("https://example.com", "GET", nil);

	fmt.Println(s.GetUnparsedHTML())
}
