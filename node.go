package astpatt

import "github.com/tehsphinx/astrav"

// Parent defines a pattern node that can have children
type Parent interface {
	CreateSubNode(nodeType astrav.NodeType) *Node
}

// Node implements a node definition
type Node struct {
	parentNode
}

// Match checks if given node matches the criteria
func (s *Node) Match(node astrav.Node) bool {
	if !s.correctType(node) {
		return false
	}
	return s.Children.Match(node)
}

func (s *Node) correctType(node astrav.Node) bool {
	return node.IsNodeType(s.NodeType)
}

type parentNode struct {
	Children nodes
	NodeType astrav.NodeType
}

// CreateSubNode creates a new node definition, appends it as a child and returns it to be adjusted.
func (s *parentNode) CreateSubNode(nodeType astrav.NodeType) *Node {
	node := &Node{
		parentNode: parentNode{
			NodeType: nodeType,
		},
	}
	s.Children = append(s.Children, node)
	return node
}
