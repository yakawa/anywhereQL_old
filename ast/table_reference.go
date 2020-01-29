package ast

/* 6.3 Table Reference */

/*
<table reference> ::=
   <table name> [ [ AS ] <correlation name> [ <left paren> <derived column list> <right paren> ] ]
 | <derived table> [ AS ] <correlation name> [ <left paren> <derived column list> <right paren> ]
 | <joined table>
<derived table> ::= <table subquery>
<derived column list> ::= <column name list>
<column name list> ::= <column name> [ { <comma> <column name> }... ]
*/

// <table reference> ::=
//   <table name> [ [ AS ] <correlation name> [ <left paren> <derived column list> <right paren> ] ]
//  | <derived table> [ AS ] <correlation name> [ <left paren> <derived column list> <right paren> ]
//  | <joined table>
type TableReference struct {
	TableName         *TableName
	DerivedTable      *DerivedTable
	CorrelationName   *string
	DerivedColumnList *DerivedColumnList
	JoinedTable       *JoinedTable
}

func (tr *TableReference) referenceNode() {}
func (tr *TableReference) Debug() string  { return "Table Reference" }

// <derived table> ::= <table subquery>
type DerivedTable struct {
	TableSubquery TableSubquery
}

func (dt *DerivedTable) referenceNode() {}
func (dt *DerivedTable) Debug() string  { return "Derived Table" }

// <derived column list> ::= <column name list>
type DerivedColumnList struct {
	ColumnNameList ColumnNameList
}

func (dc *DerivedColumnList) referenceNode() {}
func (dc *DerivedColumnList) Debug() string  { return "Derived Column List" }

// <column name list> ::= <column name> [ { <comma> <column name> }... ]
type ColumnNameList struct {
	ColumnName []string
}

func (cn *ColumnNameList) referenceNode() {}
func (cn *ColumnNameList) Debug() string  { return "Column Name List" }
