package ast

/* 8.2 Search Condition */
/*
<search condition> ::= <boolean term> | <search condition> OR <boolean term>
<boolean term> ::= <boolean factor> | <boolean term> AND <boolean factor>
<boolean factor> ::= [ NOT ] <boolean test>
<boolean test> ::= <boolean primary> [ IS [ NOT ] <truth value> ]
<truth value> ::= TRUE | FALSE | UNKNOWN
<boolean primary> ::= <predicate> | <left paren> <search condition> <right paren>
*/
// See Other <predicate>

// <search condition> ::= <boolean term> | <search condition> OR <boolean term>
type SearchCondition struct {
	BooleanTerm     *BooleanTerm
	SearchCondition *SearchCondition
}

func (sc *SearchCondition) expressionNode() {}
func (sc *SearchCondition) Debug() string   { return "SearchCondifiton" }

// <boolean term> ::= [ NOT ] <boolean test> | <boolean term> AND <boolean factor>
type BooleanTerm struct {
	IsNot         bool
	BooleanTest   *BooleanTest
	BooleanTerm   *BooleanTerm
	BooleanFactor *BooleanFactor
}

func (bt *BooleanTerm) expressionNode() {}
func (bt *BooleanTerm) Debug() string   { return "Boolean Term" }

// <boolean test> ::= {<predicate> | <left paren> <search condition> <right paren>} [ IS [ NOT ] {TRUE | FALSE | UNKNOWN} ]
type BooleanTest struct {
	IsNot           bool
	IsTrue          bool
	IsFalse         bool
	IsUnknown       bool
	Predicate       *Predicate
	SearchCondition *SearchCondition
}

func (bt *BoolenaTest) expressionNode() {}
func (bt *BoolenaTest) Debug() string   { return "Boolean Test" }
