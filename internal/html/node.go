package html

import "golang.org/x/net/html"

type HTMLAttribute struct {
	Name, Val string
}

type HTMLNodeWrapper struct {
	Prev, Next, FirstChild, LastChild, Parent *HTMLNodeWrapper

	Data  string
	Attrs []HTMLAttribute
	Type  html.NodeType // thougth it would be useless to have my own
}

func copyTree(dirtyNodeTree *html.Node) *HTMLNodeWrapper {
	if dirtyNodeTree == nil {
		return nil
	}

	node := &HTMLNodeWrapper{}
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

func (node *HTMLNodeWrapper) InsertSiblingAfter(sibling *HTMLNodeWrapper) {
	if node == nil || sibling == nil {
		return
	}

	sibling.Prev = node
	if node.Next == nil {
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

func (node *HTMLNodeWrapper) InsertSiblingBefore(sibling *HTMLNodeWrapper) {
	if node == nil || sibling == nil {
		return
	}

	sibling.Next = node
	if node.Prev == nil {
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

func (node *HTMLNodeWrapper) DeleteNode() {
	if node == nil {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		node.Parent.FirstChild = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		node.Parent.LastChild = node.Prev
	}

	node.Next = nil
	node.Prev = nil
	node.Parent = nil
}

func (node *HTMLNodeWrapper) AppendNode(child *HTMLNodeWrapper) {
	if node == nil || child == nil {
		return
	}

	if node.LastChild != nil {
		node.LastChild.Next = child
		child.Prev = node.LastChild
	} else {
		node.FirstChild = child
	}

	node.LastChild = child
	child.Parent = node
}

func (node *HTMLNodeWrapper) GetDepth() int64 {
	if node == nil {
		return -1
	}

	var depth int64 = 0

	if node.Parent == nil {
		return depth
	}

	for p := node.Parent; p != nil; p = p.Parent {
		depth++
	}

	return depth
}

func (node *HTMLNodeWrapper) Sibling() []*HTMLNodeWrapper {
	if node == nil {
		return nil
	}

	var siblings []*HTMLNodeWrapper

	// walking to first child
	var start *HTMLNodeWrapper
	if node.Parent != nil {
		start = node.Parent.FirstChild
	} else {
		for s := node.Prev; s != nil; s = s.Prev {
			start = s
		}
	}

	for s := start; s != nil; s = s.Next {
		siblings = append(siblings, s)
	}

	return siblings
}

func (node *HTMLNodeWrapper) Children() []*HTMLNodeWrapper {
	if node == nil {
		return nil
	}
	if node.FirstChild == nil {
		return nil
	}

	return node.FirstChild.Sibling()
}
