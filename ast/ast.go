package ast

import (
	"github.com/anywhereQL/anywhereQL/token"
	//"fmt"
)

type Node interface {
	Debug() string
}

type Statement interface {
	Node
	statementNode()
}

type Clause interface {
	Node
	clauseNode()
}

type Expression interface {
	Node
	expressionNode()
}

type SQL struct {
	ParentSELECTStatement SELECTStatement
	//OrderByClause
}

func (s *SQL) statementNode() {}
func (s *SQL) Debug() string  { return "" }

type SELECTStatement struct {
	SelectClause SELECTClause
}

func (ss *SELECTStatement) statementNode() {}
func (ss *SELECTStatement) Debug() string  { return "Select Statement" }

type SELECTClause struct {
	IsDistinct    bool
	IsAll         bool
	Asterisk      bool
	SelectColumns []SelectColumn
}

func (sc *SELECTClause) clauseNode()   {}
func (sc *SELECTClause) Debug() string { return "SELECT Clause" }

type SelectColumn struct {
	Name            string
	ValueExpression *ValueExpression
	AliasName       *string
	TableName       *TableName
}

type TableName struct {
	Name     string
	IsModule bool
	Catalog  *token.Token
	Schema   *token.Token
	Table    *token.Token
}

type ValueExpression struct {
}

func (ve *ValueExpression) expressionNode() {}
func (ve *ValueExpression) Debug() string   { return "" }
