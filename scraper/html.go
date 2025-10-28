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
func traverseDOM(doc *html.Node, dUnit DataUnit) {

	for n := range doc.Descendants() {

		if n.Type != html.TextNode && n.Parent.Type != html.ElementNode {
			continue
		}

		if len(dUnit.selectors["element"]) != 0 {
			for _, elem := range dUnit.selectors["element"] {
				if n.Parent.Data == elem {
					fmt.Println(n.Data)
				}
			}
		}

		if len(dUnit.selectors["class"]) != 0 {
			if n.Type == html.TextNode {
				for _, class := range dUnit.selectors["class"] {
					for _, a := range n.Parent.Attr {
						if a.Key == "class" && a.Val == class {
							fmt.Println(n.Data)
						}
					}
				}
			}
		}

	}

	// recursive search ( visits all children )
	for c := doc.FirstChild; c != nil; c = doc.NextSibling {
		traverseDOM(c, dUnit)
	}

}
