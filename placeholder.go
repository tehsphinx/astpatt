package astpatt

import "github.com/tehsphinx/astrav"

// Additional placeholder node types
const (
	NodeTypeAny  astrav.NodeType = "any"
	NodeTypeOmit astrav.NodeType = "omit"
)

// PlaceHolder is a placeholder node that can stand for none to multiple nodes.
type PlaceHolder struct {
	parentNode
}

// Populate populates the pattern node from a given ast node.
func (s *PlaceHolder) Populate(ast astrav.Node) {
	s.NodeType = NodeTypeAny
}

// Omit is a placeholder node that is to be omitted.
type Omit struct {
	parentNode
}

// Populate populates the pattern node from a given ast node.
func (s *Omit) Populate(ast astrav.Node) {
	s.NodeType = NodeTypeOmit
}
