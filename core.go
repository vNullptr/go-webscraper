package main

import (
	"webscraper/scraper"
)

func main() {

	s := scraper.NewScraper()
	s.DataUnit("main text", "string",map[string][]string{
		"class": {"price_color"},
	})
	s.DlUrl("https://books.toscrape.com", "GET", nil)
	//fmt.Println(string(s.GetUnparsedHTML()))
	doc , _ := s.ParseHTML()
	s.SearchHTML(doc)
	s.DebugShowData()

}
