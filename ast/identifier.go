package ast

/* 5.4 Names and Iddentifiers */
/*
<identifier> ::= [ <introducer><character set specification> ] <actual identifier>
<actual identifier> ::= <regular identifier> | <delimited identifier>
<SQL language identifier> ::= <SQL language identifier start> [ { <underscore> | <SQL language identifier part> }... ]
<SQL language identifier start> ::= <simple Latin letter>
<SQL language identifier part> ::= <simple Latin letter> | <digit>

<authorization identifier> ::= <identifier>
<table name> ::= <qualified name> | <qualified local table name>
<qualified local table name> ::= MODULE <period> <local table name>
<local table name> ::= <qualified identifier>
<domain name> ::= <qualified name>
<schema name> ::= [ <catalog name> <period> ] <unqualified schema name>
<unqualified schema name> ::= <identifier>
<catalog name> ::= <identifier>
<qualified name> ::= [ <schema name> <period> ] <qualified identifier>
<qualified identifier> ::= <identifier>
<column name> ::= <identifier>
<correlation name> ::= <identifier>
<module name> ::= <identifier>
<cursor name> ::= <identifier>
<procedure name> ::= <identifier>

<SQL statement name> ::= <statement name> | <extended statement name>
<statement name> ::= <identifier>
<extended statement name> ::= [ <scope option> ] <simple value specification>
<dynamic cursor name> ::= <cursor name> | <extended cursor name>
<extended cursor name> ::= [ <scope option> ] <simple value specification>
<descriptor name> ::= [ <scope option> ] <simple value specification>
<scope option> ::= GLOBAL | LOCAL
<parameter name> ::= <colon> <identifier>
<constraint name> ::= <qualified name>
<collation name> ::= <qualified name>
<character set name> ::= [ <schema name> <period> ] <SQL language identifier>
<translation name> ::= <qualified name>
<form-of-use conversion name> ::= <qualified name>
<connection name> ::= <simple value specification>
<SQL-server name> ::= <simple value specification>
<user name> ::= <simple value specification>
*/

// 5.4 <table name> ::= <qualified name> | <qualified local table name>
//     <qualified name> ::= [ <schema name> <period> ] <qualified identifier>
//     <qualified identifier> ::= <identifier>
//     <schema name> ::= [ <catalog name> <period> ] <unqualified schema name>
//     <unqualified schema name> ::= <identifier>
//     <catalog name> ::= <identifier>
//     <qualified local table name> ::= MODULE <period> <local table name>
//     <local table name> ::= <qualified identifier>
//     <qualified identifier> ::= <identifier>
// i.e. <table name> ::=
//         [[ <CatalogName> <period>] <SchemaName> <period>] <TableName>
//       | MODULE <period> <TableName>
type TableName struct {
	TableName    string
	SchemaName   *string
	CatalogName  *string
	IsLocalTable bool
}

func (tn *TableName) identifierNode() {}
func (tn *TableName) NodeLiteral() string {
	name := ""
	if tn.CatalogName != nil {
		name = name + tn.CatalogName + "."
	}
	if tn.SchemaName != nil {
		name = name + tn.SchemaName + "."
	}
	name = name + TableName
	if IsLocalTable == true {
		name = "Local Table Name: " + name
	} else {
		name = "Table Name: " + name
	}
	return name
}
