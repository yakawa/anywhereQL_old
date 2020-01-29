package ast

/* 7.5 Joined Table */
/*
<joined table> ::= <cross join> | <qualified join> | <left paren> <joined table> <right paren>
<cross join> ::= <table reference> CROSS JOIN <table reference>
<qualified join> ::= <table reference> [ NATURAL ] [ <join type> ] JOIN <table reference> [ <join specification> ]
<join specification> ::= <join condition> | <named columns join>
<join condition> ::= ON <search condition>
<named columns join> ::= USING <left paren> <join column list> <right paren>
<join type> ::= INNER | <outer join type> [ OUTER ] | UNION
<outer join type> ::= LEFT | RIGHT | FULL
<join column list> ::= <column name list>
*/
// See Other <table reference>, <column name list>, <search condition>

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
