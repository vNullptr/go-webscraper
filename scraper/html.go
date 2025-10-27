package scraper

import (
	"bytes"
	"fmt"

	"golang.org/x/net/html"
)

func (s *Scraper) ParseHTML() error {

	if len(s.targetData) == 0 {
		panic("Target data units not initialized !") 
	}

	rawHTML := bytes.NewReader(s.GetUnparsedHTML())
	doc, err := html.Parse(rawHTML)
	if err != nil {
		fmt.Println("error happened while parsing html")
		return fmt.Errorf("error happened while parsing html : %w", err)
	}

	// idk if i want to do this -> PS 2
	for _, v := range s.targetData {
		traverseDOM(doc, v)
	} 

	return nil

}

// PS : this is a test still thinking about how to make this "modular" to take any combination of selectors
// PS 2 : decided on how to have custom selectors and store the found data, still deciding on how to do the search 
func traverseDOM(doc *html.Node, selector DataUnit){

	for n := range doc.Descendants() {
		if ( n.Type == html.TextNode && (n.Parent.Type == html.ElementNode && n.Parent.Data == selector.selectors["element"][0]) ) {
			fmt.Println(n.Data)
		}
	}

	// recursive search ( visits all children )
	for c := doc.FirstChild; c != nil; c = doc.NextSibling {
		traverseDOM(c, selector)
	}
}