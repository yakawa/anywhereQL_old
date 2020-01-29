package ast

/* 7.11 scalar subquery /  row subquery / table subquery / subquery
/*
<scalar subquery> ::= <subquery>
<row subquery> ::= <subquery>
<table subquery> ::= <subquery>
<subquery> ::= <left paren> <query expression> <right paren>
*/

// <scalar subquery> ::= <subquery>
type ScalarSubquery struct {
	Subquery Subquery
}

func (ss *ScalarSubquery) expressionNode() {}
func (ss *ScalarSubquery) Debug() string   { return "Scalar Subquery" }

// <row subquery> ::= <subquery>
type RowSubquery struct {
	Subquery Subquery
}

func (rs *RowSubquery) expressionNode() {}
func (rs *RowSubquery) Debug() string   { return "Row Subquery" }

// <table subquery> ::= <subquery>
type TableSubquery struct {
	Subquery Subquery
}

func (ts *TableSubquery) expressionNode() {}
func (ts *TableSubquery) Debug() string   { return "Table Subquery" }

// <subquery> ::= <left paren> <query expression> <right paren>
type Subquery struct {
	QueryExpression QueryExpression
}

func (s *Subquery) expressionNode() {}
func (s *Subquery) Debug() string   { return "Subquery" }
