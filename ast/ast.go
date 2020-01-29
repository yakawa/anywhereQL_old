package ast

import (
	"fmt"
)

type Node interface {
	Debug() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Term interface {
	Node
	termNode()
}

type Primary interface {
	Node
	primaryNode()
}

type Spec interface {
	Node
	specNode()
}

type Identifier interface {
	Node
	identifierNode()
}
