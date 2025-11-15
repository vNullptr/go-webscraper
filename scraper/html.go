package scraper

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func (s *Scraper) ParseHTML() error {

	if len(s.targetData) == 0 {
		panic("Target data units not initialized !");
	}

	rawHTML := bytes.NewReader(s.GetUnparsedHTML());
	doc, err := html.Parse(rawHTML);
	if err != nil {
		fmt.Println("error happened while parsing html");
		return fmt.Errorf("error happened while parsing html : %w", err);
	}

	s.htmlRoot = doc;

	return nil;

}

// might seem useless until i make the a node tree to store cleaned up html.Node tree
// basically these are wrappers
func (s *Scraper) DOM() *html.Node {
	return s.htmlRoot;
}

func (s *Scraper) FirstChild(node *html.Node) *html.Node {
	if node == nil { return nil }
	return node.FirstChild;
}

func (s *Scraper) NthChild(node *html.Node, index int) *html.Node {

	if node == nil { return nil }

	child := node;

	for range index {

		if node.NextSibling == nil {
			return nil;
		}
		child = node.NextSibling;

	}

	return child;
}

func (s *Scraper) Parent(node *html.Node) *html.Node {
	if node == nil { return nil; }
	return node.Parent;
}

func (s *Scraper) Children(node *html.Node) []*html.Node {
	if node == nil { return nil; }
	
	var childrens []*html.Node;
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		childrens = append(childrens, c);
	}

	return childrens;
}

func (s *Scraper) GetAttr(node *html.Node) []html.Attribute {
	if node == nil { return nil; }
	return node.Attr;
}

func (s *Scraper) HasAttr(node *html.Node, key string, value string) bool {
	if node == nil { return false }

	attr := s.GetAttr(node)
	for index := range attr {
		if attr[index].Key == key {
			if strings.Contains(attr[index].Val, value) {
				return true;
			}
		}
	}
	return false;
}

// wrapper for clarity uses HasAttr()
func (s *Scraper) HasClass(node *html.Node, class string ) bool {
	return s.HasAttr(node, "class", class);
}

// will add error handling later ( winApi style with lastError in the struct )
func (s *Scraper) FirstByAttr(node *html.Node, key string, value string) *html.Node {

  	if node == nil { return nil; }

	if s.HasAttr(node, key, value) {
		return node;
	}

	for c := node; c != nil; c = c.NextSibling {
		return s.FirstByAttr(c, key, value);
	}

	return nil;
}

func (s *Scraper) IsTag(node *html.Node, tagname string) bool {
	
	if node == nil { return false; }
	
	if node.Type == html.ElementNode {
		if node.Data == tagname {
			return true;
		}
	}

	return false;
}

// will still keep the other commonly used function and keep this one for more specific search
func (s *Scraper) WalkToFirst(node *html.Node, callback func(*html.Node)bool) *html.Node{

	if node == nil { return nil }

	if callback(node){
		return node;
	}

	for c := node; c != nil; c = c.NextSibling {
		return s.WalkToFirst(node, callback);
	}

	return nil;

}

func (s *Scraper) Ancestors(node *html.Node) []*html.Node{

	if node == nil { return nil }

	var ancestors []*html.Node;

	for c := node.Parent; c != nil; c = c.Parent {
		ancestors = append(ancestors, c);
	}

	return ancestors;
}