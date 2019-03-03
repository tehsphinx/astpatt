package astpatt

import "github.com/tehsphinx/astrav"

// AssignStmt node
type AssignStmt struct {
	parentNode
}

// Populate populates the pattern node from a given ast node.
func (s *AssignStmt) Populate(ast astrav.Node) {
	s.populateDefault(ast)
}

// FuncDecl node
type FuncDecl struct {
	parentNode
	Name string
}

// Populate populates the pattern node from a given ast node.
func (s *FuncDecl) Populate(ast astrav.Node) {
	s.populateDefault(ast)
	if named, ok := ast.(astrav.Named); ok {
		s.Name = named.NodeName().Name
	}

	s.populateChildren(ast)
}

// IfStmt node
type IfStmt struct {
	parentNode
}

// Populate populates the pattern node from a given ast node.
func (s *IfStmt) Populate(ast astrav.Node) {
	s.populateDefault(ast)

	for _, ast := range ast.Children() {
		if !ast.IsNodeType(astrav.NodeTypeBlockStmt) {
			continue
		}

		node := creator(ast)
		s.Nodes = append(s.Nodes, node)
	}
}

// ReturnStmt node
type ReturnStmt struct {
	parentNode
}

// Populate populates the pattern node from a given ast node.
func (s *ReturnStmt) Populate(ast astrav.Node) {
	s.populateDefault(ast)
}
