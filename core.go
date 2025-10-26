package main

import (
	"webscraper/scraper"
)


func main() {

	s := scraper.NewScraper()
	s.DlUrl("https://www.example.com", "GET", nil)
	//fmt.Println(string(s.GetUnparsedHTML()))
	s.ParseHTML()

}
