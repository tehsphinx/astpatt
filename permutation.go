package astpatt

import (
	"github.com/tehsphinx/astrav"
)

// permutations creates a bunch of possible pattern permutations with the same meaning.
func permutations(ast astrav.Node) []Node {
	childPermuts := permutChildren(ast)
	childCombis := combinations(childPermuts)

	var nodes []Node
	if len(childCombis) == 0 {
		node := createNode(ast.NodeType())
		node.fillNode(nil, ast)
		nodes = append(nodes, node)
	}
	for _, perm := range childCombis {
		permutations := permuteNodes(perm, ast)

		for _, p := range permutations {
			permNode := createNode(ast.NodeType())
			permNode.fillNode(p, ast)
			nodes = append(nodes, permNode)
		}
	}
	return nodes
}

func permutChildren(ast astrav.Node) [][]Node {
	var childPermuts [][]Node
	for _, child := range ast.Children() {
		var (
			children []Node
			nn       = createNode(child.NodeType())
		)
		if nn.isType(NodeTypeOmit) {
			continue
		}
		if nn.isType(NodeTypeSkip) {
			chPerms := permutChildren(child)
			childPermuts = append(childPermuts, chPerms...)
			continue
		} else {
			children = permutations(child)
		}
		childPermuts = append(childPermuts, children)
	}
	return childPermuts
}

func permuteNodes(nodes []Node, parentAST astrav.Node) [][]Node {
	var perms [][]Node
	if parentAST.NodeType() == astrav.NodeTypeFile {
		perms = Perm(nodes)
	} else {
		perms = append(perms, nodes)
	}
	return perms
}

func combinations(pNodes [][]Node) (combinations [][]Node) {
	comb := func(input []Node, prev [][]Node) (output [][]Node) {
		if len(prev) == 0 {
			prev = append(prev, []Node{})
		}

		for _, node := range input {
			for _, prevNode := range prev {
				output = append(output, append(prevNode, node))
			}
		}
		return output
	}

	for _, nodes := range pNodes {
		combinations = comb(nodes, combinations)
	}
	return combinations
}

// Perm calls f with each permutation of a.
func Perm(nodes []Node) [][]Node {
	return perm(nodes, nil, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []Node, nodes [][]Node, i int) [][]Node {
	if i > len(a) {
		var b []Node
		return append(nodes, append(b, a...))
	}
	nodes = perm(a, nodes, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		nodes = perm(a, nodes, i+1)
		a[i], a[j] = a[j], a[i]
	}
	return nodes
}
