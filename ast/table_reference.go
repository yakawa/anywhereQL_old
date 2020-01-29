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

type TableReference struct {
	Table   TableName
	Name    *string
	Columns *DerivedColumnList
}

func (tr *TableReference) referenceNode() {}
func (tr *TableReference) Debug() string  { return "Table Reference" }
