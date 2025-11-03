package scraper

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func (s *Scraper) ParseHTML() ( error ) {

	if len(s.targetData) == 0 {
		panic("Target data units not initialized !")
	}

	rawHTML := bytes.NewReader(s.GetUnparsedHTML())
	doc, err := html.Parse(rawHTML)
	if err != nil {
		fmt.Println("error happened while parsing html")
		return fmt.Errorf("error happened while parsing html : %w", err)
	}

	s.htmlRoot = doc

	return nil

}


// might seem useless until i make the a node tree to store cleaned up html.Node tree 
// basically these are wrappers
func (s *Scraper) DOM() *html.Node{
	return s.htmlRoot
}

func (s *Scraper) FirstChild(node *html.Node) *html.Node {
	return node.FirstChild
}

func (s *Scraper) NthChild( node *html.Node, index int) *html.Node {
	child := node

	for range index {

		if node.NextSibling == nil { 
			return nil 
		}
		child = node.NextSibling

	}

	return child 
}

func (s *Scraper) Parent(node *html.Node) *html.Node {
	return node.Parent
}

func (s *Scraper) Children(node *html.Node) []*html.Node {
	var childrens []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		childrens = append(childrens, c)
	}

	return childrens
}


//	These functions below will be deleted at some point


func (s *Scraper) SearchHTML() {
	for i := range s.targetData {
		s.traverseDOM(s.htmlRoot, i)
	}
}

// PS : decided to go with this simple / unefficient structure for now 
// until i have enough time to make my own node tree to clean up the html.Node mess to make the search better n faster
func (s *Scraper) traverseDOM(doc *html.Node, index int) {

	for n := range doc.Descendants() {

		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "href" {
					// will start a Scrape on another thread
				}
			}
		}

		if n.Type != html.TextNode && n.Parent.Type != html.ElementNode {
			continue
		}

		if len(s.targetData[index].selectors["element"]) != 0 {
			for _, elem := range s.targetData[index].selectors["element"] {
				if n.Parent.Type == html.ElementNode && n.Parent.Data == elem {
					// might want to make it return instead of directly store 
					s.targetData[index].data = append(s.targetData[index].data, n.Data)
				}
			}
		}

		if len(s.targetData[index].selectors["class"]) != 0 {
			if n.Type == html.TextNode {
				for _, class := range s.targetData[index].selectors["class"] {
					for _, a := range n.Parent.Attr {
						if a.Key == "class" && strings.Contains(a.Val, class) {
							s.targetData[index].data = append(s.targetData[index].data, n.Data)
						}
					}
				}
			}
		}

		if len(s.targetData[index].selectors["id"]) != 0 {
			if n.Type == html.TextNode {
				for _, id := range s.targetData[index].selectors["id"] {
					for _, a := range n.Parent.Attr {
						if a.Key == "id" && a.Val == id {
							s.targetData[index].data = append(s.targetData[index].data, n.Data)
						}
					}
				}
			}
		}

	}

	// recursive search ( visits all children )
	for c := doc.FirstChild; c != nil; c = doc.NextSibling {
		s.traverseDOM(c, index)
	}

}
