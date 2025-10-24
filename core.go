package main

import (
	"fmt"
	"webscraper/scraper"
)

func main() {

	s := scraper.Scraper{}
	body, _ := s.DlUrl("https://www.example.com", "GET")
	fmt.Println(string(body))

}
