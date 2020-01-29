package ast

/* 7. Query Expression */
/*
# 7.1 Row Value Constructor
<row value constructor> ::= <row value constructor element> | <left paren> <row value constructor list> <right paren> | <row subquery>
<row value constructor list> ::= <row value constructor element> [ { <comma> <row value constructor element> }... ]
<row value constructor element> ::= <value expression> | <null specification> | <default specification>
<null specification> ::= NULL
<default specification> ::= DEFAULT

# 7.2  Table Value Constructor
<table value constructor> ::= VALUES <table value constructor list>
<table value constructor list> ::= <row value constructor> [ { <comma> <row value constructor> }... ]

# 7.3 Table Expression
<table expression> ::= <from clause> [ <where clause> ] [ <group by clause> ] [ <having clause> ]

# 7.4 From Clause
<from clause> ::= FROM <table reference> [ { <comma> <table reference> }... ]

# 7.5 Joined Table
<joined table> ::= <cross join> | <qualified join> | <left paren> <joined table> <right paren>
<cross join> ::= <table reference> CROSS JOIN <table reference>
<qualified join> ::= <table reference> [ NATURAL ] [ <join type> ] JOIN <table reference> [ <join specification> ]
<join specification> ::= <join condition> | <named columns join>
<join condition> ::= ON <search condition>
<named columns join> ::= USING <left paren> <join column list> <right paren>
<join type> ::= INNER | <outer join type> [ OUTER ] | UNION
<outer join type> ::= LEFT | RIGHT | FULL
<join column list> ::= <column name list>

# 7.6 Where Clause
<where clause> ::= WHERE <search condition>

# 7.7 Group By Clause
<group by clause> ::= GROUP BY <grouping column reference list>
<grouping column reference list> ::= <grouping column reference> [ { <comma> <grouping column reference> }... ]
<grouping column reference> ::= <column reference> [ <collate clause> ]

# 7.8 Having Clause
<having clause> ::= HAVING <search condition>

# 7.9 Query Specification
<query specification> ::= SELECT [ <set quantifier> ] <select list> <table expression>
<select list> ::= <asterisk> | <select sublist> [ { <comma> <select sublist> }... ]
<select sublist> ::= <derived column> | <qualifier> <period> <asterisk>
<derived column> ::= <value expression> [ <as clause> ]
<as clause> ::= [ AS ] <column name>

# 7.10 Query Expression
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

# 7.11 Scalar SubQuery / Row SubQuery / Table SubQuery
<scalar subquery> ::= <subquery>
<row subquery> ::= <subquery>
<table subquery> ::= <subquery>
<subquery> ::= <left paren> <query expression> <right paren>
*/

// <row value constructor> ::= <row value constructor element> | <left paren> <row value constructor list> <right paren> | <row subquery>
type RowValueConstructor struct {
	RowValueConstructorElement *RowValueConstructorElement
	RowValueConstructorList    *RowValueConstructorList
	RowSubquery                *RowSubquery
}

func (rc *RowValueConstructor) expressionNode() {}
func (rc *RowValueConstructor) Debug() string   { return "Row Value Constructor" }

// <row value constructor list> ::= <row value constructor element> [ { <comma> <row value constructor element> }... ]
type RowValueConstructorList struct {
	RowValueConstructorElement []RowValueConstructorElement
}

func (rc *RowValueConstructorList) expressionNode() {}
func (rc *RowValueConstructorList) Debug() string   { return "RowValueConstructorList" }

// <row value constructor element> ::= <value expression> | NULL | DEFAULT
type RowValueConstructorElement struct {
	ValueExpression *ValueExpression
	IsNull          bool
	IsDefault       bool
}

func (re *RowValueConstructorElement) expressionNode() {}
func (re *RowValueConstructorElement) Debug() string   { return "Row Value Constructor Element" }

// <table value constructor> ::= VALUES <table value constructor list>
type TableValueConstructor struct {
	TableValueConstructorList TableValueConstructorList
}

func (tc *TableValueConstructor) expressionNode() {}
func (tc *TableValueConstructor) Debug() string   { return "Table Value Constructor" }

// <table value constructor list> ::= <row value constructor> [ { <comma> <row value constructor> }... ]
type TableValueConstructorList struct {
	RowValueConstructor []RowValueConstructor
}

func (tc *TableValueConstructorList) expressionNode() {}
func (tc *TableValueConstructorList) Debug() string   { return "Table Value Constructor List" }

// <table expression> ::= <from clause> [ <where clause> ] [ <group by clause> ] [ <having clause> ]
type TableExpression struct {
	FromClause    FromClause
	WhereClause   *WhereClause
	GroupByClause *GroupByClase
	HavingClause  *HavingClause
}

func (te *TableExpression) expressionNode() {}
func (te *TableExpression) Debug() string   { return "Table Expression" }

// <from clause> ::= FROM <table reference> [ { <comma> <table reference> }... ]
type FromClause struct {
	TableReference []TableReference
}

func (fc *FromClause) expressionNode() {}
func (fc *FromClause) Debug() string   { return "From Clause" }

// <joined table> ::= <cross join> | <qualified join> | <left paren> <joined table> <right paren>
type JoinedTable struct {
	CrossJoin     *CrossJoin
	QualifiedJoin *QualifiedJoin
	JoinedTable   *JoinedTable
}

func (jt *JoinedTable) identifierNode() {}
func (jt *JoinedTable) Debug() string   { return "Joined Table" }

// <cross join> ::= <table reference> CROSS JOIN <table reference>
type CrossJoin struct {
	LeftTableReference  TableReference
	RightTableReference TableReference
}

func (cj *CrossJoin) identifierNode() {}
func (cj *CrossJoin) Debug() string   { return "Cross Join" }

// <qualified join> ::= <table reference> [ NATURAL ] [ {INNER | {LEFT | RIGHT | FULL} [OUTER] | UNION} ] JOIN <table reference> [ <join specification> ]
// <join specification> ::= ON <search condigion> | USING <left paren> <column name list> <right paren>
type QualifiedJoin struct {
	LeftTableReference  TableReference
	RightTableReference TableReference
	SearchCondition     *SearchCondition
	ColumnNameList      *ColumnNameList
	IsNatural           bool
	IsInner             bool
	IsLeft              bool
	IsRight             bool
	IsFull              bool
	IsUnion             bool
}

// <where clause> ::= WHERE <search condition>
type WhereClause struct {
	SearchCondition SearchCondition
}

func (wc *WhereClause) expressionNode() {}
func (wc *WhereClause) Debug() string   { return "Where Clause" }

// <group by clause> ::= GROUP BY <grouping column reference list>
// <grouping column reference list> ::= <grouping column reference> [ { <comma> <grouping column reference> }... ]
type GroupByClause struct {
	GroupingColumnReference []GroupingColumnReference
}

func (gc *GroupByClause) expressionNode() {}
func (gc *GroupByClause) Debug() string   { return "Group By Clause" }

// <grouping column reference> ::= <column reference> [ <collate clause> ]
type GroupingColumnReference struct {
	ColumnReference ColumnReference
	CollateClause   *CollateClause
}

func (gr *GroupingColumnReference) expressionNode() {}
func (gr *GroupingColumnReference) Debug() string   { return "Grouping Column reference" }

// <having clause> ::= HAVING <search condition>
type HavingClause struct {
	SearchCondition SearchCondition
}

func (hc *HavingClause) expressionNode() {}
func (hc *HavingClause) Debug() string   { return "Having Clause" }

// <query specification> ::= SELECT [ DISTINCT | ALL ] <select list> <table expression>
type QuerySpecification struct {
	IsDistinct      bool
	IsAll           bool
	SelectList      SelectList
	TableExpression TableExpression
}

func (qs *QuerySpecification) expressionNode() {}
func (qs *QuerySpecification) Debug() string   { return "QuerySpecification" }

// <select list> ::= <asterisk> | <select sublist> [ { <comma> <select sublist> }... ]
type SelectList struct {
	IsAsterisk    bool
	SelectSublist []SelectSublist
}

func (sl *SelectList) expressionNode() {}
func (sl *SelectList) Debug() string   { return "Select List" }

// <select sublist> ::= <derived column> | <qualifier> <period> <asterisk>
type SelectSublist struct {
	DerivedColumn *DerivedColumn
	Qualifier     *Qualifier
}

func (ss *SelectSublist) expressionNode() {}
func (ss *SelectSublist) Debug() string   { return "Select Sublist" }

// <derived column> ::= <value expression> [ [ AS ] <column name> ]
type DerivedColumn struct {
	ValueExpression ValueExpression
	ColumnName      *string
}

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
