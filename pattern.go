package astpatt

import "github.com/tehsphinx/astrav"

// Matcher represents criterias describing an ast criteria
type Matcher interface {
	Match(node astrav.Node) bool
}

// NewPattern creates a new pattern
func NewPattern() *Pattern {
	return &Pattern{}
}

// Pattern implements a solutions pattern
type Pattern struct {
	parentNode
}

// Match checks the pattern against given parent node
func (s *Pattern) Match(parent astrav.Node) bool {
	return s.Children.Match(parent)
}
