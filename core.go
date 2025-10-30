package main

import (
	"webscraper/scraper"
)

func main() {

	s := scraper.NewScraper()
	s.DataUnit("main text", "string",map[string][]string{
		"element": {"h2", "h1"},
	})
	s.DlUrl("https://www.coursera.org", "GET", nil)
	//fmt.Println(string(s.GetUnparsedHTML()))
	doc , _ := s.ParseHTML()
	s.SearchHTML(doc)
	s.DebugShowData()

}
