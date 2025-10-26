package scraper

import (
	"bytes"
	"fmt"

	"golang.org/x/net/html"
)

func (s *Scraper) ParseHTML() {

	rawHTML := bytes.NewReader(s.GetUnparsedHTML())
	doc, err := html.Parse(rawHTML)
	if err != nil {
		fmt.Println("error happened while parsing html")
	}

	traverseDOM(doc, "p")

}

// PS : this is a test still thinking about how to make this "modular" to take any combination of selectors
func traverseDOM(doc *html.Node, elementSelector string){

	for n := range doc.Descendants() {
		if ( n.Type == html.TextNode && (n.Parent.Type == html.ElementNode && n.Parent.Data == elementSelector) ) {
			fmt.Println(n.Data)
		}
	}

	// recursive search ( visits all children )
	for c := doc.FirstChild; c != nil; c = doc.NextSibling {
		traverseDOM(c, elementSelector)
	}
}