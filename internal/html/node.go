package html

import "golang.org/x/net/html"

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

	node := &HTMLNodeWpr{}
	node.Data = dirtyNodeTree.Data
	node.Type = dirtyNodeTree.Type

	if len(dirtyNodeTree.Attr) > 0 {
		for _, attr := range dirtyNodeTree.Attr {
			newAttr := HTMLAttribute{}
			newAttr.Name = attr.Key
			newAttr.Val = attr.Val
			node.Attrs = append(node.Attrs, newAttr)
		}
	}

	// childrens
	if dirtyNodeTree.FirstChild != nil {
		child := copyTree(dirtyNodeTree.FirstChild)
		if node.FirstChild == nil { 
			node.FirstChild = child 
		}
		node.LastChild = child
		node.AppendNode(child)
	}
	
	//siblings
	if dirtyNodeTree.NextSibling != nil {
		sibling := copyTree(dirtyNodeTree.NextSibling)
		node.InsertSiblingAfter(sibling)
	}
	

	return node
}

func (node *HTMLNodeWpr) InsertSiblingAfter(sibling *HTMLNodeWpr){
	if (node == nil || sibling == nil) { return }
	
	sibling.Prev = node
	if (node.Next == nil){
		node.Parent.LastChild = sibling
	} else {
		sibling.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = sibling
	}

	node.Next = sibling
	sibling.Parent = node.Parent

}

func (node *HTMLNodeWpr) InsertSiblingBefore(sibling *HTMLNodeWpr){
    if node == nil || sibling == nil { return }
    
	sibling.Next = node
	if (node.Prev == nil){
		node.Parent.FirstChild = sibling
	} else {
		sibling.Prev = node.Prev
	}
    
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

func (node *HTMLNodeWpr) GetDepth() int {
	if node == nil {return -1}

	if node.Parent == nil {
		return 0
	}

	return 1 + node.Parent.GetDepth()
}

func (node *HTMLNodeWpr) Sibling() []*HTMLNodeWpr{
	if node == nil {return nil}
	
	var siblings []*HTMLNodeWpr
	for s := node; s != nil; s = s.Next {
		siblings = append(siblings, s)
	}

	return siblings
}

func (node *HTMLNodeWpr) Children() []*HTMLNodeWpr{
	if node == nil {return nil}
	if node.FirstChild == nil {return nil}

	return node.FirstChild.Sibling()
}