package astpatt

import "github.com/tehsphinx/astrav"

// Special node types
const (
	// NodeTypeOmit will be omitted including its children.
	NodeTypeOmit astrav.NodeType = "omit"
	// NodeTypeSkip defines a node that is not taken into accound. It's children will be inlined.
	NodeTypeSkip astrav.NodeType = "skip"
)

// Omit is a placeholder node that is to be omitted.
type Omit struct {
	parentNode
}

// Populate populates the pattern node from a given ast node.
func (s *Omit) Populate(ast astrav.Node) {
	s.NodeType = NodeTypeOmit
}

// Skip is a placeholder node that is to be omitted.
type Skip struct {
	parentNode
}

// Populate populates the pattern node from a given ast node.
func (s *Skip) Populate(ast astrav.Node) {
	s.NodeType = NodeTypeSkip
}
