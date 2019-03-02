package astpatt

import "github.com/tehsphinx/astrav"

type nodes []*Node

// Match checks if given node matches the criteria
func (s nodes) Match(parent astrav.Node) bool {
	nodes := parent.Children()
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
