package scraper

import "golang.org/x/net/html"

// no idea if i should make structs for attribute and type
// instead of using a string map
type HTMLAttribute struct {
	Name, Val string
}

type HTMLNodeWpr struct {

	Prev, Next, FirstChild, LastChild, Parent *HTMLNodeWpr

	Data string
	Attrs []HTMLAttribute
	Type html.NodeType // thougth it would be useless to have my own
}

func copyTree(dirtyNodeTree *html.Node) *HTMLNodeWpr {
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
	sibling.Parent = node.Parent
}

func (node *HTMLNodeWpr) InsertSiblingBefore(sibling *HTMLNodeWpr){
    if node == nil || sibling == nil { return }
    
	sibling.Next = node
    sibling.Prev = node.Prev
    
	if node.Prev != nil {
        node.Prev.Next = sibling
    }
    
	node.Prev = sibling
	sibling.Parent = node.Parent
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
	node.Parent = nil
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
	child.Parent = node
}

func (node *HTMLNodeWpr) HasChild() bool {
	if node == nil { return false }

	return node.FirstChild != nil

}

func (node *HTMLNodeWpr) GetDepth() int {
	if node == nil {return -1}

	if node.Parent == nil {
		return 0
	}

	return 1 + node.Parent.GetDepth()
}
