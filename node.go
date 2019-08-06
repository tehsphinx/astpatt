package astpatt

import (
	"github.com/tehsphinx/astrav"
)

// Node represents a pattern node.
type Node interface {
	Match(Node) bool
	Populate(astrav.Node)
	Children() []Node
	Walk(f func(node Node) bool)

	isType(astrav.NodeType) bool
	populateDefault(ast astrav.Node)
	fillNode(nodes []Node, ast astrav.Node)
}

// DefaultNode implements the default node definition
type DefaultNode struct {
	parentNode
}

type parentNode struct {
	Nodes    nodes           `json:"children,omitempty"`
	NodeType astrav.NodeType `json:"type"`
	Code     string          `json:"-"`
}

func (s *parentNode) correctType(node Node) bool {
	return node.isType(s.NodeType)
}

// Match checks if given node matches the criteria
func (s *parentNode) Match(node Node) bool {
	if !s.correctType(node) {
		return false
	}
	return s.Nodes.Match(node)
}

// Children returns the child nodes.
func (s *parentNode) Children() []Node {
	return s.Nodes
}

// Populate populates the pattern node from a given ast node.
// This is the default Populator implementation. Overwrite for spefic pattern node behavior.
func (s *parentNode) Populate(ast astrav.Node) {
	s.populateDefault(ast)
	s.populateChildren(ast)
}

func (s *parentNode) populateDefault(ast astrav.Node) {
	s.NodeType = ast.NodeType()
	s.Code = ast.GetSourceString()
}

func (s *parentNode) fillNode(nodes []Node, ast astrav.Node) {
	s.Nodes = nodes
	s.populateDefault(ast)
}

func (s *parentNode) populateChildren(ast astrav.Node) {
	for _, ast := range ast.Children() {
		node := creator(ast)
		if node.isType(NodeTypeOmit) {
			continue
		}
		if node.isType(NodeTypeSkip) {
			s.populateChildren(ast)
			continue
		}
		s.Nodes = append(s.Nodes, node)
	}
}

// Walk traverses the tree and its children.
// return false to skip children of the current element
func (s *parentNode) Walk(f func(node Node) bool) {
	if !f(s) {
		return
	}

	for _, child := range s.Children() {
		child.Walk(f)
	}
}

func (s *parentNode) populateBlock(ast astrav.Node) {
	for _, ast := range ast.Children() {
		if !ast.IsNodeType(astrav.NodeTypeBlockStmt) {
			continue
		}
		node := creator(ast)
		s.Nodes = append(s.Nodes, node)
	}
}

func (s *parentNode) isType(nodeType astrav.NodeType) bool {
	return s.NodeType == nodeType
}

type nodes []Node

// Match checks if given node matches the criteria
func (s nodes) Match(parent Node) bool {
	var nodes = parent.Children()
	if len(s) != len(nodes) {
		return false
	}
	for i, child := range s {
		if !child.Match(nodes[i]) {
			return false
		}
	}
	return true
}
