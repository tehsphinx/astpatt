package astpatt

import (
	"github.com/tehsphinx/astrav"
)

// FuncDecl node
type FuncDecl struct {
	parentNode
	Name string
}

// Populate populates the pattern node from a given ast node.
func (s *FuncDecl) Populate(ast astrav.Node) {
	s.populateDefault(ast)
	if named, ok := ast.(astrav.Named); ok {
		s.Name = named.NodeName()
	}

	s.populateChildren(ast)
}

func (s *FuncDecl) fillNode(nodes []Node, ast astrav.Node) {
	s.Nodes = nodes
	s.populateDefault(ast)
	if named, ok := ast.(astrav.Named); ok {
		s.Name = named.NodeName()
	}
}

// SelectorExpr node
type SelectorExpr struct {
	parentNode
	Name string
}

// Populate populates the pattern node from a given ast node.
func (s *SelectorExpr) Populate(ast astrav.Node) {
	s.populateDefault(ast)
	if named, ok := ast.(astrav.Named); ok {
		s.Name = named.NodeName()
	}

	s.populateChildren(ast)
}

func (s *SelectorExpr) fillNode(nodes []Node, ast astrav.Node) {
	s.Nodes = nodes
	s.populateDefault(ast)
	if named, ok := ast.(astrav.Named); ok {
		s.Name = named.NodeName()
	}
}

// Match checks if given node matches the criteria
func (s *SelectorExpr) Match(node Node) bool {
	if !s.correctType(node) {
		return false
	}
	if n, ok := node.(*SelectorExpr); ok && n.Name != s.Name {
		return false
	}
	return s.Nodes.Match(node)
}
