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
