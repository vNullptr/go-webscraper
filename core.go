package main

import (
	"webscraper/scraper"
)

func main() {

	s := scraper.NewScraper()
	s.AppendDataUnit("main text", map[string][]string{
		"element": {"h2", "h1"},
	})
	s.DlUrl("https://www.coursera.org", "GET", nil)
	//fmt.Println(string(s.GetUnparsedHTML()))
	s.ParseHTML()

}
