package ast

/* 7.10 Query Expression */
/*
<query expression> ::= <non-join query expression> | <joined table>
<non-join query expression> ::=
  <non-join query term>
 | <query expression> UNION  [ ALL ] [ <corresponding spec> ] <query term>
 | <query expression> EXCEPT [ ALL ] [ <corresponding spec> ] <query term>
<query term> ::= <non-join query term> | <joined table>
<non-join query term> ::= <non-join query primary> | <query term> INTERSECT [ ALL ] [ <corresponding spec> ] <query primary>
<query primary> ::= <non-join query primary> | <joined table>
<non-join query primary> ::= <simple table> | <left paren> <non-join query expression> <right paren>
<simple table> ::= <query specification> | <table value constructor> | <explicit table>
<explicit table> ::= TABLE <table name>
<corresponding spec> ::= CORRESPONDING [ BY <left paren> <corresponding column list> <right paren> ]
<corresponding column list> ::= <column name list>
*/
// See Other <joined table>, <query specification>, <table value constructor>, <table name>, <column name list>

// <query expression> ::= <non-join query expression> | <joined table>
type QueryExpression struct {
	NonJoinQueryExpression *NonJoinQueryExpression
	JoinedTable            *JoinedTable
}

func (qe *QueryExpression) expressionNode() {}
func (qe *QueryExpression) Debug() string   { return "Query Expression" }

// <non-join query expression> ::= <non-join query term> | <query expression> {UNION | EXCEPT} [ ALL ] [ <corresponding spec> ] <query term>
type NonJoinQueryExpression struct {
	NonJoinQueryTerm  *NonJoinQueryTem
	QueryExpression   *QueryExpression
	IsUnion           bool
	IsExcept          bool
	IsAll             bool
	CorrespondingSpec *CorrespondingSpec
	QueryTerm         *QueryTerm
}

func (ne *NonJoinQueryExpression) expressionNode() {}
func (ne *NonJoinQueryExpression) Debug() string   { return "Non Join Query Expression" }

// <query term> ::= <non-join query term> | <joined table>
type QueryTerm struct {
	NonJoinQueryTerm *NonJoinQueryTerm
	JoinedTable      *JoinedTable
}

func (qt *QueryTerm) termNode()     {}
func (qt *QueryTerm) Debug() string { return "Query Term" }

// <non-join query term> ::= <non-join query primary> | <query term> INTERSECT [ ALL ] [ <corresponding spec> ] <query primary>
type NonJoinQueryTerm struct {
	NonJoinQueryPrimary *NonJoinQueryPrimary
	IsIntersect         bool
	IsAll               bool
	QueryTerm           *QueryTerm
	CorrespondingSpec   *CorrespondingSpec
	QueryPrimary        *QueryPrimary
}

func (nt *NonJoinQueryTerm) termNode()     {}
func (nt *NonJoinQueryTerm) Debug() string { return "Non Join Query Term" }

// <query primary> ::= <non-join query primary> | <joined table>
type QueryPrimary struct {
	NonJoinQueryPrimary *NonJoinQueryPrimary
	JoinedTable         *JoinedTable
}

func (qp *QueryPrimary) primaryNode()  {}
func (qp *QueryPrimary) Debug() string { return "Query Primary" }

// <non-join query primary> ::= <simple table> | <left paren> <non-join query expression> <right paren>
type NonJoinQueryPrimary struct {
	SimpleTable            *SimpleTable
	NonJoinQueryExpression *NonJoinQueryExpression
}

func (np *NonJoinQueryPrimary) primaryNode() {}
func (np *NonJoinQueryPrimary) Debug()       { return "Non Join Query Primary" }

// <simple table> ::= <query specification> | <table value constructor> | TABLE <table name>
type SimpleTable struct {
	QuerySpecification    *QuerySpecification
	TableValueConstructor *TableValueConstructor
	TableName             *TableName
}

func (st *SimpleTable) identifierNode() {}
func (st *SimpleTable) Debug() string   { return "Simple Table" }

// <corresponding spec> ::= CORRESPONDING [ BY <left paren> <corresponding column list> <right paren> ]
type CorrespondingSpec struct {
	ColumnNameList *ColumnNameList
}

func (cs *CorrespondingSpec) specNode()     {}
func (cs *CorrespondingSpec) Debug() string { return "Corresponding Spec" }
