package astpatt

import (
	"github.com/tehsphinx/astrav"
)

func creator(astNode astrav.Node) Node {
	nodeType := astNode.NodeType()

	var n Node
	switch nodeType {
	case astrav.NodeTypeComment:
		n = &Omit{}
	case astrav.NodeTypeCommentGroup:
		n = &Omit{}
	// case astrav.NodeTypeField:
	// 	n = &Field{}
	// case astrav.NodeTypeFieldList:
	// 	n = &FieldList{}
	// case astrav.NodeTypeBadExpr:
	// 	n = &BadExpr{}
	case astrav.NodeTypeIdent:
		n = &Omit{}
	// case astrav.NodeTypeEllipsis:
	// 	n = &Ellipsis{}
	// case astrav.NodeTypeBasicLit:
	// 	n = &Omit{}
	// case astrav.NodeTypeFuncLit:
	// 	n = &FuncLit{}
	// case astrav.NodeTypeCompositeLit:
	// 	n = &CompositeLit{}
	// case astrav.NodeTypeParenExpr:
	// 	n = &ParenExpr{}
	case astrav.NodeTypeSelectorExpr:
		n = &SelectorExpr{}
	// case astrav.NodeTypeIndexExpr:
	// 	n = &IndexExpr{}
	// case astrav.NodeTypeSliceExpr:
	// 	n = &SliceExpr{}
	// case astrav.NodeTypeTypeAssertExpr:
	// 	n = &TypeAssertExpr{}
	// case astrav.NodeTypeCallExpr:
	// 	n = &CallExpr{}
	// case astrav.NodeTypeStarExpr:
	// 	n = &StarExpr{}
	// case astrav.NodeTypeUnaryExpr:
	// 	n = &UnaryExpr{}
	// case astrav.NodeTypeBinaryExpr:
	// 	n = &BinaryExpr{}
	// case astrav.NodeTypeKeyValueExpr:
	// 	n = &KeyValueExpr{}
	// case astrav.NodeTypeArrayType:
	// 	n = &ArrayType{}
	// case astrav.NodeTypeStructType:
	// 	n = &StructType{}
	// case astrav.NodeTypeFuncType:
	// 	n = &Omit{}
	// case astrav.NodeTypeInterfaceType:
	// 	n = &InterfaceType{}
	// case astrav.NodeTypeMapType:
	// 	n = &MapType{}
	// case astrav.NodeTypeChanType:
	// 	n = &ChanType{}
	// case astrav.NodeTypeBadStmt:
	// 	n = &BadStmt{}
	// case astrav.NodeTypeDeclStmt:
	// 	n = &Omit{}
	// case astrav.NodeTypeEmptyStmt:
	// 	n = &EmptyStmt{}
	// case astrav.NodeTypeLabeledStmt:
	// 	n = &LabeledStmt{}
	// case astrav.NodeTypeExprStmt:
	// 	n = &ExprStmt{}
	// case astrav.NodeTypeSendStmt:
	// 	n = &SendStmt{}
	// case astrav.NodeTypeIncDecStmt:
	// 	n = &Omit{}
	case astrav.NodeTypeAssignStmt:
		n = &Skip{}
	// case astrav.NodeTypeGoStmt:
	// 	n = &GoStmt{}
	// case astrav.NodeTypeDeferStmt:
	// 	n = &DeferStmt{}
	// case astrav.NodeTypeReturnStmt:
	// 	n = &ReturnStmt{}
	// case astrav.NodeTypeBranchStmt:
	// 	n = &BranchStmt{}
	// case astrav.NodeTypeBlockStmt:
	// 	n = &BlockStmt{}
	// case astrav.NodeTypeIfStmt:
	// 	n = &IfStmt{}
	// case astrav.NodeTypeCaseClause:
	// 	n = &CaseClause{}
	// case astrav.NodeTypeSwitchStmt:
	// 	n = &SwitchStmt{}
	// case astrav.NodeTypeTypeSwitchStmt:
	// 	n = &TypeSwitchStmt{}
	// case astrav.NodeTypeCommClause:
	// 	n = &CommClause{}
	// case astrav.NodeTypeSelectStmt:
	// 	n = &SelectStmt{}
	// case astrav.NodeTypeForStmt:
	// 	n = &ForStmt{}
	// case astrav.NodeTypeRangeStmt:
	// 	n = &RangeStmt{}
	// case astrav.NodeTypeImportSpec:
	// 	n = &ImportSpec{}
	// case astrav.NodeTypeValueSpec:
	// 	n = &ValueSpec{}
	// case astrav.NodeTypeTypeSpec:
	// 	n = &TypeSpec{}
	// case astrav.NodeTypeBadDecl:
	// 	n = &BadDecl{}
	// case astrav.NodeTypeGenDecl:
	// 	n = &Omit{}
	case astrav.NodeTypeFuncDecl:
		n = &FuncDecl{}
	// case astrav.NodeTypeFile:
	// 	n = &File{}
	case astrav.NodeTypePackage:
		n = &Pattern{}
	default:
		n = &DefaultNode{}
	}

	n.Populate(astNode)
	return n
}
