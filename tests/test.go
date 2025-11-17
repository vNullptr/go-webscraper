package main

import (
	"fmt"

	scraper "webscraper/internal/scraper"
)


func main() {

	s := scraper.NewScraper();
	s.FetchURL("https://example.com", "GET", nil);
	s.ParseHTML();
	
	// fmt.Println(string(s.GetUnparsedHTML()));
	
	// not practical for now, but will be a lot easier to use when im done with the cleanedup node tree
	root := s.DOM();
	head := s.FirstChild(s.NthChild(root, 1))
	body := head.NextSibling
	a := s.Children(s.Children(body)[0])[0].NextSibling.NextSibling.FirstChild
	fmt.Println(s.GetAttr(a)) 

}
