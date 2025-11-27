package scraper

import "golang.org/x/net/html"

// no idea if i should make structs for attribute and type
// instead of using a string map
type HTMLAttribute struct {
	Name, Val string
}

type HTMLNodeWpr struct {

	Prev, Next, FirstChild, LastChild *HTMLNodeWpr

	Data string
	Attrs []HTMLAttribute
	Type html.NodeType // thougth it would be useless to have my own
}

func cleanUp(dirtyNodeTree *html.Node) *HTMLNodeWpr {
	if dirtyNodeTree == nil {
		return nil
	}

	node := HTMLNodeWpr{}

	return &node
}

func (node *HTMLNodeWpr) InsertSiblingAfter(sibling *HTMLNodeWpr){
	if (node == nil || sibling == nil) { return }
	
	sibling.Prev = node
	sibling.Next = node.Next

	if node.Next != nil {
		node.Next.Prev = sibling
	}

	node.Next = sibling
}

func (node *HTMLNodeWpr) InsertSiblingBefore(sibling *HTMLNodeWpr){
    if node == nil || sibling == nil { return }
    
	sibling.Next = node
    sibling.Prev = node.Prev
    
	if node.Prev != nil {
        node.Prev.Next = sibling
    }
    
	node.Prev = sibling
}

func (node *HTMLNodeWpr) DeleteNode(){
	if node == nil { return }

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Next = nil
	node.Prev = nil 
}

func (node *HTMLNodeWpr) AppendNode(child *HTMLNodeWpr){
	if node == nil || child == nil {return}

	if (node.LastChild != nil){
		node.LastChild.Next = child
		child.Prev = node.LastChild
	} else {
		node.FirstChild = child
	}

	node.LastChild = child
}

func (node *HTMLNodeWpr) HasChild() bool {
	if node == nil {return false}

	if node.FirstChild == nil && node.LastChild == nil {
		return false 
	} 

	return true
}


