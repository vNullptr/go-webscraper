package scraper

import (
	"bytes"
	"fmt"
	"strings"

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

	s.htmlRoot = doc

	return nil

}

// might seem useless until i make the a node tree to store cleaned up html.Node tree
// basically these are wrappers
func (s *Scraper) DOM() *html.Node {
	return s.htmlRoot
}

func (s *Scraper) FirstChild(node *html.Node) *html.Node {
	return node.FirstChild
}

func (s *Scraper) NthChild(node *html.Node, index int) *html.Node {
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

func (s *Scraper) GetAttr(node *html.Node) []html.Attribute {
	return node.Attr
}

func (s *Scraper) HasAttr(node *html.Node, key string, value string) bool {
	attr := s.GetAttr(node)
	for index := range attr {
		if attr[index].Key == key {
			if strings.Contains(attr[index].Val, value) {
				return true
			}
		}
	}
	return false
}

func (s *Scraper) FindByAttr(root *html.Node, key string, value string) *html.Node {

	if s.HasAttr(root, key, value) {
		return root
	}

	for c := root; c != nil; c = c.NextSibling {
		return s.FindByAttr(c, key, value)
	}

	return nil
}

func (s *Scraper) IsTag(node *html.Node, tagname string) bool {
	if node.Type == html.ElementNode {
		if node.Data == tagname {
			return true 
		}
	}

	return false
}

// this search's tags but can't be doing one for class by tag then class by class ect ect ...
// will probably use a func callback as argument
func (s *Scraper) FindUntilTag(root *html.Node, key string, value string, endKey string) *html.Node {
	
	if s.HasAttr(root, key, value) {
		return root
	}

	if s.IsTag(root, endKey) {
		return nil
	}

	for c := root; c != nil; c = c.NextSibling {
		return s.FindUntilTag(c, key, value, endKey)
	}

	return nil
}
