package ast

// 8.1 / 8.2 / 8.3 / 8.4 / 8.5 / 8.6 / 8.7 / 8.8 / 8.9 / 8.10 / 8.11 Predicates
/*
<predicate> ::=
  <comparison predicate>
 | <between predicate>
 | <in predicate>
 | <like predicate>
 | <null predicate>
 | <quantified comparison predicate>
 | <exists predicate>
 | <unique predicate>
 | <match predicate>
 | <overlaps predicate>
<comparison predicate> ::= <row value constructor> <comp op> <row value constructor>
<comp op> ::=
  <equals operator>
 | <not equals operator>
 | <less than operator>
 | <greater than operator>
 | <less than or equals operator>
 | <greater than or equals operator>
<between predicate> ::= <row value constructor> [ NOT ] BETWEEN <row value constructor> AND <row value constructor>
<in predicate> ::= <row value constructor> [ NOT ] IN <in predicate value>
<in predicate value> ::= <table subquery> | <left paren> <in value list> <right paren>
<in value list> ::= <value expression> { <comma> <value expression> }...
<like predicate> ::= <match value> [ NOT ] LIKE <pattern> [ ESCAPE <escape character> ]
<match value> ::= <character value expression>
<pattern> ::= <character value expression>
<escape character> ::= <character value expression>
<null predicate> ::= <row value constructor> IS [ NOT ] NULL
<quantified comparison predicate> ::= <row value constructor> <comp op> <quantifier> <table subquery>
<quantifier> ::= <all> | <some>
<all> ::= ALL
<some> ::= SOME | ANY
<exists predicate> ::= EXISTS <table subquery>
<unique predicate> ::= UNIQUE <table subquery>
<match predicate> ::= <row value constructor> MATCH [ UNIQUE ] [ PARTIAL | FULL ] <table subquery>
<overlaps predicate> ::= <row value constructor 1> OVERLAPS <row value constructor 2>
<row value constructor 1> ::= <row value constructor>
<row value constructor 2> ::= <row value constructor>
*/
// See <table subquery>

// <predicate> ::=
//  <comparison predicate>
// | <between predicate>
// | <in predicate>
// | <like predicate>
// | <null predicate>
// | <quantified comparison predicate>
// | <exists predicate>
// | <unique predicate>
// | <match predicate>
// | <overlaps predicate>
type Predicate struct {
	ComparisonPredicate           *ComparisonPredicate
	BetweenPredicate              *BetweenPredicate
	InPredicate                   *InPredicate
	LikePredicate                 *LikePredicate
	NullPredicate                 *NullPredicate
	QuantifiedComparisonPredicate *QuantifiedComparisonPredicate
	ExistsPredicate               *ExistsPredicate
	UniquePredicate               *UniquePredicate
	MatchPredicate                *MatchPredicate
	OverlapsPredicate             *OverlapsPredicate
}

func (p *Predicate) expressionNode() {}
func (p *Predicate) Debug() string   { return "Predicate" }

// <comparison predicate> ::= <row value constructor> <comp op> <row value constructor>
// <comp op> ::=
//  <equals operator>
//  | <not equals operator>
//  | <less than operator>
//  | <greater than operator>
//  | <less than or equals operator>
//  | <greater than or equals operator>
type ComparisonPredicate struct {
	Equals               bool
	NotEquals            bool
	LessThan             bool
	GreaterThan          bool
	LessThanEquals       bool
	GreaterThanEquals    bool
	RowValueConstructor1 RowValueConstructor
	RowValueConstructor2 RowValueConstructor
}

func (cp *ComparisonPredicate) expressionNode() {}
func (cp *ComparisonPredicate) Debug() string   { return "ComparisonPredicate" }

// <between predicate> ::= <row value constructor> [ NOT ] BETWEEN <row value constructor> AND <row value constructor>
type BetweenPredicate struct {
	IsNot      bool
	RowValue   RowValueConstructor
	RangeBegin RowValueConstructor
	RangeEnd   RowValueConstructor
}

func (bp *BetweenPredicate) expressionNode() {}
func (bp *BetweenPredicate) Debug() string   { return "Between Predicate" }

// <in predicate> ::= <row value constructor> [ NOT ] IN <in predicate value>
// <in predicate value> ::= <table subquery> | <left paren> <in value list> <right paren>
// <in value list> ::= <value expression> { <comma> <value expression> }...
type InPredicate struct {
	IsNot           bool
	RowValue        RowValueConstructor
	TableSubquery   *TableSubquery
	ValueExpression []ValueExpression
}

func (ip *InPredicate) expressionNode() {}
func (ip *InPredicate) Debug() string   { return "InPredicate" }

// <like predicate> ::= <match value> [ NOT ] LIKE <pattern> [ ESCAPE <escape character> ]
// <match value> ::= <character value expression>
// <pattern> ::= <character value expression>
// <escape character> ::= <character value expression>
type LikePredicate struct {
	MatchValue CharacterValueExpression
	IsNot      bool
	Pattern    CharacterValueExpression
	Escape     *CharacterValueExpression
}

func (lp *LikePredicate) expressionNode() {}
func (lp *LikePredicate) Debug() string   { return "Like Predicate" }

// <null predicate> ::= <row value constructor> IS [ NOT ] NULL
type NullPredicate struct {
	RowValue RowValueConstructor
	IsNot    bool
}

func (np *NullPredicate) expressionNode() {}
func (np *NullPredicate) Debug() string   { return "Null Predicate" }

// <quantified comparison predicate> ::= <row value constructor> <comp op> {ALL | SOME | ANY} <table subquery>
type QuantifiedComparisonPredicate struct {
	Equals               bool
	NotEquals            bool
	LessThan             bool
	GreaterThan          bool
	LessThanEquals       bool
	GreaterThanEquals    bool
	RowValueConstructor1 RowValueConstructor
	TableSubquery        TableSubquery
}

func (qp *QuantifiedComparisonPredicate) expressionNode() {}
func (qp *QuantifiedComparisonPredicate) Debug() string   { return "QuantifiedComparisonPredicate" }

// <exists predicate> ::= EXISTS <table subquery>
type ExistsPredicate struct {
	TableSubquery TableSubquery
}

func (ep *ExistsPredicate) expressionNode() {}
func (ep *ExistsPredicate) Debug() string   { return "ExistsPredicate" }

// <unique predicate> ::= UNIQUE <table subquery>
type UniquePredicate struct {
	TableSubquery TableSubquery
}

func (ep *UniquePredicate) expressionNode() {}
func (ep *UniquePredicate) Debug() string   { return "UniquePredicate" }

// <match predicate> ::= <row value constructor> MATCH [ UNIQUE ] [ PARTIAL | FULL ] <table subquery>
type MatchPredicate struct {
	RowValue      RowValueConstructor
	IsUnique      bool
	IsPartial     bool
	IsFull        bool
	TableSubquery TableSubquery
}

func (mp *MatchPredicate) expressionNode() {}
func (mp *MatchPredicate) Debug() string   { return "Match Predicate" }

// <overlaps predicate> ::= <row value constructor 1> OVERLAPS <row value constructor 2>
// <row value constructor 1> ::= <row value constructor>
// <row value constructor 2> ::= <row value constructor>
type OverlapsPredicate struct {
	RowValue1 RowValueConstructor
	RowValue2 RowValueConstructor
}

func (op *OverlapsPredicate) expressionNode() {}
func (op *OverlapsPredicate) Debug() string   { return "Overlaps Predicate" }
