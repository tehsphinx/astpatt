package astpatt

import "github.com/tehsphinx/astrav"

// ExtractPattern extracts the pattern from a given ast package
func ExtractPattern(pkg *astrav.Package) *Pattern {
	pattern := NewPattern()
	pattern.Populate(pkg)
	return pattern
}

// NewPattern creates a new pattern
func NewPattern() *Pattern {
	return &Pattern{}
}

// Pattern implements a solutions pattern
type Pattern struct {
	parentNode
}

// MatchPkg checks the pattern against given parent node
func (s *Pattern) MatchPkg(pkg *astrav.Package) bool {
	pkgPatt := ExtractPattern(pkg)
	return s.Match(pkgPatt)
}

// Match checks the pattern against another pattern
func (s *Pattern) Match(node Node) bool {
	return s.Nodes.Match(node)
}
