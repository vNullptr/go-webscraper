package scraper

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func (s *Scraper) ParseHTML() ( *html.Node, error) {

	if len(s.targetData) == 0 {
		panic("Target data units not initialized !")
	}

	rawHTML := bytes.NewReader(s.GetUnparsedHTML())
	doc, err := html.Parse(rawHTML)
	if err != nil {
		fmt.Println("error happened while parsing html")
		return nil,fmt.Errorf("error happened while parsing html : %w", err)
	}

	return doc,nil

}


func (s *Scraper) SearchHTML( doc *html.Node ) {
	for i := range s.targetData {
		traverseDOM(doc, &s.targetData[i])
	}
}

// PS : decided to go with this simple / unefficient structure for now 
// until i have enough time to make my own node tree to clean up the html.Node mess to make the search better n faster
func traverseDOM(doc *html.Node, dUnit *DataUnit) {

	for n := range doc.Descendants() {

		if n.Type != html.TextNode && n.Parent.Type != html.ElementNode {
			continue
		}

		if len(dUnit.selectors["element"]) != 0 {
			for _, elem := range dUnit.selectors["element"] {
				if n.Parent.Type == html.ElementNode && n.Parent.Data == elem {
					// might want to make it return instead of directly store 
					dUnit.data = append(dUnit.data, n.Data)
				}
			}
		}

		if len(dUnit.selectors["class"]) != 0 {
			if n.Type == html.TextNode {
				for _, class := range dUnit.selectors["class"] {
					for _, a := range n.Parent.Attr {
						if a.Key == "class" && strings.Contains(a.Val, class) {
							dUnit.data = append(dUnit.data, n.Data)
						}
					}
				}
			}
		}

		if len(dUnit.selectors["id"]) != 0 {
			if n.Type == html.TextNode {
				for _, id := range dUnit.selectors["id"] {
					for _, a := range n.Parent.Attr {
						if a.Key == "id" && a.Val == id {
							dUnit.data = append(dUnit.data, n.Data)
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
