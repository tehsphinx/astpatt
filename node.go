package astpatt

import "github.com/tehsphinx/astrav"

// Node represents a pattern node.
type Node interface {
	Match(Node) bool
	Populate(astrav.Node)
	Children() []Node

	isType(astrav.NodeType) bool
}

// DefaultNode implements the default node definition
type DefaultNode struct {
	parentNode
}

type parentNode struct {
	Nodes    nodes           `json:"children,omitempty"`
	NodeType astrav.NodeType `json:"type"`
	Code     string          `json:"code,omitempty"`
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

func (s *parentNode) populateChildren(ast astrav.Node) {
	var prevNode Node
	for _, ast := range ast.Children() {
		node := creator(ast)
		if node.isType(NodeTypeOmit) {
			continue
		}
		if node.isType(NodeTypeAny) && prevNode != nil && prevNode.isType(NodeTypeAny) {
			continue
		}
		s.Nodes = append(s.Nodes, node)
		prevNode = node
	}
}

func (s *parentNode) isType(nodeType astrav.NodeType) bool {
	return s.NodeType == nodeType
}

type nodes []Node

// Match checks if given node matches the criteria
func (s nodes) Match(parent Node) bool {
	var (
		nodeIndex int
		nodes     = parent.Children()
	)
	for _, child := range s {
		if !child.Match(nodes[nodeIndex]) {
			return false
		}
		nodeIndex++
	}
	return true
}
