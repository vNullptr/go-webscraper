package scraper

import "golang.org/x/net/html"

// no idea if i should make structs for attribute and type
// instead of using a string map
type HTMLAttribute struct {
	Name, Val string
}

type HTMLType uint16

type HTMLNodeWpr struct {

	Prev, Next *HTMLNodeWpr

	Data string
	Attr HTMLAttribute
	Type HTMLType 
}

func cleanUp(dirtyNodeTree *html.Node) HTMLNodeWpr {

	node := HTMLNodeWpr{}

	return node

}

func (node *HTMLNodeWpr) InsertSiblingAfter(sibling *HTMLNodeWpr){
	if node.Next != nil {
		sibling.Next = node.Next
	}

	node.Next = sibling
}

func (node *HTMLNodeWpr) InsertSiblingBefore(sibling *HTMLNodeWpr){
	if node.Prev!= nil {
		sibling.Prev = node.Prev
	}

	node.Prev = sibling
}
