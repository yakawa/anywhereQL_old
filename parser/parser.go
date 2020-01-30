package parser

import (
	"github.com/anywhereQL/anywhereQL/ast"
	"github.com/anywhereQL/anywhereQL/token"
)

type Parser struct {
	tokens []token.Token
	pos    int
}

func New(tokens []token.Token) *Parser {
	p := &Parser{tokens: tokens, pos: 0}
	return p
}

func (p *Parser) getToken() token.Token {
	if p.pos >= len(p.tokens) {
		return token.Token{Type: token.EOF_TOKEN, Literal: ""}
	}
	return p.tokens[p.pos]
}

func (p *Parser) getNextToken() token.Token {
	if p.pos+1 >= len(p.tokens) {
		return token.Token{Type: token.EOF_TOKEN, Literal: ""}
	}
	p.pos += 1
	return p.tokens[p.pos]
}

func (p *Parser) peekToken(n int) token.Token {
	if p.pos+n >= len(p.tokens) {
		return token.Token{Type: token.EOF_TOKEN, Literal: ""}
	}
	return p.tokens[p.pos+n]
}

func (p *Parser) GetSQL() *ast.SQL {
	sql := &ast.SQL{}
	t := p.getToken()
	for t.Type != token.EOF_TOKEN {
		switch t.Type {
		case token.KEYWORD_SELECT_TOKEN:
			p.parseSelectStatement()
		}
		t = p.getToken()
	}
	return sql
}

func (p *Parser) parseSelectStatement() *ast.SELECTStatement {
	// <query specification> ::= SELECT [ <set quantifier> ] <select list> <from clause> [ <where clause> ] [ <group by clause> ] [ <having clause> ]
	r := &ast.SELECTStatement{}
	s := p.parseSelectClause()
	if s != nil {
		r.SelectClause = *s
	}
	return r
}

func (p *Parser) parseSelectClause() *ast.SELECTClause {
	// <SELECT Clause> ::= [ DISTINCT | ALL ] { <asterisk> | <select sublist> [ { <comma> <select sublist> }... ] }

	r := &ast.SELECTClause{
		Asterisk:   false,
		IsDistinct: false,
		IsAll:      false,
	}
	t := p.getNextToken()
BL:
	for t.Type != token.KEYWORD_FROM_TOKEN && t.Type != token.EOF_TOKEN {
		switch t.Type {
		case token.KEYWORD_DISTINCT_TOKEN:
			r.IsDistinct = true
		case token.KEYWORD_ALL_TOKEN:
			r.IsAll = true
		case token.ASTERISK_TOKEN:
			if len(r.SelectColumns) != 0 {
				return nil
			}
			r.Asterisk = true
		case token.COMMA_TOKEN:
			if r.Asterisk == true {
				return nil
			}
		default:
			lineToken := make([]token.Token, 0)
			for !(t.Type == token.EOF_TOKEN || t.Type == token.COMMA_TOKEN || t.Type == token.KEYWORD_FROM_TOKEN) {
				lineToken = append(lineToken, t)
				t = p.getNextToken()
			}
			c := p.parseSelectColumn(lineToken)
			if c == nil {
				return nil
			}
			r.SelectColumns = append(r.SelectColumns, *c)
			if t.Type == token.EOF_TOKEN || t.Type == token.KEYWORD_FROM_TOKEN {
				break BL
			}
		}
		t = p.getNextToken()
	}
	return r
}

func (p *Parser) parseSelectColumn(t []token.Token) *ast.SelectColumn {
	// <select sublist> ::= <value expression> [ [ AS ] <column name> ] | <qualifier> <period> <asterisk>
	// <qualifier> ::= <table name> | <correlation name>
	// <table name> ::= <qualified name> | <qualified local table name>
	// <qualified local table name> ::= MODULE <period> <local table name>
	// <local table name> ::= <qualified identifier>
	// <qualified name> ::= [ <schema name> <period> ] <qualified identifier>
	// <schema name> ::= [ <catalog name> <period> ] <unqualified schema name>
	// <qualified identifier> ::= <identifier>
	// <unqualified schema name> ::= <identifier>
	// <catalog name> ::= <identifier>
	// <correlation name> ::= <identifier>
	//  <column name> ::= <identifier>
	//
	// | <value expression> [ [ AS ] <identifier> ]
	// | <identifier> <period> <asterisk>
	// | <identifier> <period> <identifier> <period> <asterisk>
	// | <identifier> <period> <identifier> <period> <identifier> <period> <asterisk>
	// | MODULE <period> <identifier> <period> <asterisk>

	if t[len(t)-1].Type == token.ASTERISK_TOKEN {
		p.parseQualifier(t)
	} else {
		p.parseValueExpression(t)
	}
	return nil
}

func (p *Parser) parseQualifier(t []token.Token) *ast.TableName {
	// | <identifier> <period> <asterisk>
	// | <identifier> <period> <identifier> <period> <asterisk>
	// | <identifier> <period> <identifier> <period> <identifier> <period> <asterisk>
	// | MODULE <period> <identifier> <period> <asterisk>
	tbl := &ast.TableName{Name: "", Catalog: nil, Schema: nil, Table: nil, IsModule: false}
	if t[0].Type == token.KEYWORD_MODULE_TOKEN {
		// | MODULE <period> <identifier> <period> <asterisk>
		if len(t) != 5 {
			return nil
		}
		if t[1].Type == token.PERIOD_TOKEN && t[3].Type == token.PERIOD_TOKEN && t[4].Type == token.ASTERISK_TOKEN {
			if t[2].Type != token.IDENTIFIER_TOKEN {
				return nil
			}
			tbl.IsModule = true
			tbl.Name = t[2].Literal
			tbl.Table = &t[2]
			return tbl
		}
		return nil
	}
	if len(t) == 3 {
		// | <identifier> <period> <asterisk>
		if t[0].Type == token.IDENTIFIER_TOKEN && t[1].Type == token.PERIOD_TOKEN && t[2].Type == token.ASTERISK_TOKEN {
			tbl.Table = &t[0]
			tbl.Name = t[0].Literal
			return tbl
		}
		return nil
	} else if len(t) == 5 {
		// | <identifier> <period> <identifier> <period> <asterisk>
		for i, tt := range t {
			if i == 0 || i == 2 {
				if tt.Type != token.IDENTIFIER_TOKEN {
					return nil
				}
				if i == 0 {
					tbl.Schema = &tt
					tbl.Name += tt.Literal
				} else {
					tbl.Table = &tt
					tbl.Name = tbl.Name + "." + tt.Literal
				}
			} else if i == 1 || i == 3 {
				if tt.Type != token.PERIOD_TOKEN {
					return nil
				}
			} else if i == 4 {
				if tt.Type != token.ASTERISK_TOKEN {
					return nil
				}
			}
		}
		return tbl
	} else if len(t) == 7 {
		// | <identifier> <period> <identifier> <period> <identifier> <period> <asterisk>
		for i, tt := range t {
			if i == 0 || i == 2 || i == 4 {
				if tt.Type != token.IDENTIFIER_TOKEN {
					return nil
				}
				if i == 0 {
					tbl.Schema = &tt
					tbl.Name += tt.Literal
				} else if i == 2 {
					tbl.Table = &tt
					tbl.Name = tbl.Name + "." + tt.Literal
				} else {
					tbl.Catalog = &tt
					tbl.Name = tbl.Name + "." + tt.Literal
				}
			} else if i == 1 || i == 3 || i == 5 {
				if tt.Type != token.PERIOD_TOKEN {
					return nil
				}
			} else if i == 6 {
				if tt.Type != token.ASTERISK_TOKEN {
					return nil
				}
			}
		}
		return tbl
	}
	return nil
}
func (p *Parser) parseValueExpression(t []token.Token) *ast.ValueExpression {
	// <value expression> [ [ AS ] <identifier> ]
	// <value expression> ::= <numeric value expression> | <string value expression> | <datetime value expression> | <interval value expression>
	// <numeric value expression> ::= <term> | <numeric value expression> <plus sign> <term> | <numeric value expression> <minus sign> <term>
	// <term> ::= <factor> | <term> <asterisk> <factor> | <term> <solidus> <factor>
	// <factor> ::= [ <sign> ] <numeric primary>
	// <numeric primary> ::= <value expression primary> | <numeric value function>

	// <value expression primary> ::= <unsigned value specification> | <column reference> | <set function specification> | <scalar subquery> | <case expression> | <left paren> <value expression> <right paren> | <cast specification>
	tkn = p.getToken()
	if tkn.Type == token.KEYWORD_POSITION_TOKEN || tkn.Type == token.KEYWORD_EXTRACT_TOKEN ||
		tkn.Type == token.KEYWORD_CHAR_LENGTH_TOKEN || tkn.Type == token.KEYWORD_CHARACTER_LENGTH_TOKEN ||
		tkn.Type == token.KEYWORD_OCTET_LENGTH_TOKEN || tkn.Type == token.KEYWORD_BIT_LENGTH_TOKEN {
		p.parseNumericValueFunction()
	}
	return nil
}

func (p *Parser) parseNumericValueFunction() {
	// <numeric value function> ::= <position expression> | <extract expression> | <length expression>
	// <position expression> ::= POSITION <left paren> <character value expression> IN <character value expression> <right paren>
	// <extract expression> ::= EXTRACT <left paren> <extract field> FROM <extract source> <right paren>
	// <length expression> ::= <char length expression> | <octet length expression> | <bit length expression>
	// <char length expression> ::= { CHAR_LENGTH | CHARACTER_LENGTH } <left paren> <string value expression> <right paren>
	// <octet length expression> ::= OCTET_LENGTH <left paren> <string value expression> <right paren>
	// <bit length expression> ::= BIT_LENGTH <left paren> <string value expression> <right paren>

}
