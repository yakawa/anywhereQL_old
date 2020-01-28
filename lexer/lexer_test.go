package lexer

import (
	"errors"
	"fmt"
	"testing"

	"github.com/anywhereQL/anywhereQL/token"
)

func TestReadSpecialCharacterToken(t *testing.T) {
	tests := []struct {
		input           string
		expectedToken   token.TokenType
		expectedLiteral string
	}{
		{" ", token.SPACE_TOKEN, " "},
		{"\"", token.DOUBLE_QUOTE_TOKEN, "\""},
		{"%", token.PERCENT_TOKEN, "%"},
		{"&", token.AMPERSAND_TOKEN, "&"},
		{"'", token.QUOTE_TOKEN, "'"},
		{"(", token.LEFT_PAREN_TOKEN, "("},
		{")", token.RIGHT_PAREN_TOKEN, ")"},
		{"*", token.ASTERISK_TOKEN, "*"},
		{"+", token.PLUS_SIGN_TOKEN, "+"},
		{",", token.COMMA_TOKEN, ","},
		{"-", token.MINUS_SIGN_TOKEN, "-"},
		{".", token.PERIOD_TOKEN, "."},
		{"/", token.SOLIDAS_TOKEN, "/"},
		{":", token.COLON_TOKEN, ":"},
		{";", token.SEMICOLON_TOKEN, ";"},
		{"<", token.LESS_THAN_OPERATOR_TOKEN, "<"},
		{"=", token.EQUALS_OPERATOR_TOKEN, "="},
		{">", token.GREATER_THAN_OPERATOR_TOKEN, ">"},
		{"?", token.QUESTION_MARK_TOKEN, "?"},
		{"_", token.UNDERSCORE_TOKEN, "_"},
		{"|", token.VERTICAL_BAR_TOKEN, "|"},
		{"[", token.LEFT_BRACKET_TOKEN, "["},
		{"]", token.RIGHT_BRACKET_TOKEN, "]"},
		{"<>", token.NOT_EQUALS_OPERATOR_TOKEN, "<>"},
		{">=", token.GREATER_THAN_OR_EQUALS_OPERATOR_TOKEN, ">="},
		{"<=", token.LESS_THAN_OR_EQUALS_OPERATOR_TOKEN, "<="},
		{"||", token.CONCATENATION_OPERATOR_TOKEN, "||"},
		{"..", token.DOUBLE_PERIOD_TOKEN, ".."},
	}

	for i, tt := range tests {
		l := New(tt.input)
		tok, li, err := l.readSpecialCharacterToken()
		if err != nil {
			t.Errorf("[%d] Err error is not nil: %s (%s)\n", i, err, tt.input)
		} else {
			if tok != tt.expectedToken {
				t.Errorf("[%d] Err: Token type mistmatch (expected: %s, but got %s)\n", i, tt.expectedToken, tok)
			}
			if li != tt.expectedLiteral {
				t.Errorf("[%d] Err: Token Literal mistmatch (expected: %s, but got: %s)\n", i, tt.expectedLiteral, li)
			}
		}
	}
}

func TestReadCommentToken(t *testing.T) {
	tests := []struct {
		input           string
		expectedLiteral string
	}{
		{"--- Comment\n", "Comment"},
		{"--- Comment", "Comment"},
		{"-----1+2\n", "1+2"},
		{"-- 1---2\n", "1---2"},
		{"-- Expect: None\n", "Expect: None"},
	}
	for i, tt := range tests {
		l := New(tt.input)
		s, err := l.readComment()
		if err != nil {
			t.Fatalf("[%d] Err error is not nil: %s (%s)\n", i, err, tt.input)
		} else {
			if s != tt.expectedLiteral {
				t.Errorf("[%d] Err: Token Literal mistmatch (expected: %s, but got: %s)\n", i, tt.expectedLiteral, s)
			}
		}
	}
}

func TestReadNumberToken(t *testing.T) {
	tests := []struct {
		input           string
		expectedLiteral string
		expectedNil     bool
	}{
		{"1", "1", false},
		{"123", "123", false},
		{"123.", "123.", false},
		{"123.4", "123.4", false},
		{".1", ".1", false},
		{"1.1.1", "1.1", false},
		{"1..2", "1.", false},
		{"1..", "1.", false},
		{"1E+1", "1E+1", false},
		{"1E-1", "1E-1", false},
		{"1E1", "1E1", false},
		{"12E10", "12E10", false},
		{"12.3E-2", "12.3E-2", false},
		{"1.E+2", "1.E+2", false},
		{".2E+2", ".2E+2", false},
		{"1E+", "1E+", true},
		{".E+3", ".", true},
		{"B'0101'", "B'0101'", false},
		{"B'0101' \t '0101'", "B'0101' \t '0101'", false},
		{"B'0101' -- Comment '0101'", "B'0101' -- Comment '0101'", false},
		{"B'0101' -- Comment\n '0101'", "B'0101' -- Comment\n '0101'", false},
		{"B'0201' -- Comment\n '0101'", "B'0101' -- Comment\n '0101'", true},
		{"X'0A01'", "X'0A01'", false},
		{"X'0a01' \t '0101'", "X'0a01' \t '0101'", false},
		{"X'0A01' -- Comment '0101'", "X'0A01' -- Comment '0101'", false},
		{"X'0A01' -- Comment\n '0101'", "X'0A01' -- Comment\n '0101'", false},
		{"X'0K01' -- Comment\n '0101'", "X'0101' -- Comment\n '0101'", true},
		{"B'0101' \t SELECT", "B'0101' \t ", false},
		{"X'0101' \t SELECT", "X'0101' \t ", false},
	}
	for i, tt := range tests {
		l := New(tt.input)
		s, err := l.readNumber()
		if err != nil && tt.expectedNil == false {
			t.Errorf("[%d] Err error is not nil: %s (%s)\n", i, err, tt.input)
		}
		if err == nil && tt.expectedNil == true {
			t.Errorf("[%d] Err error is nil (%s)\n", i, tt.input)
		}
		if s != tt.expectedLiteral && tt.expectedNil == false {
			t.Errorf("[%d] Err: Token Literal mistmatch (expected: %s, but got: %s)\n", i, tt.expectedLiteral, s)
		}
	}
}

func TestReadIdentifier(t *testing.T) {
	tests := []struct {
		input           string
		expectedToken   token.TokenType
		expectedLiteral string
	}{
		{"ABSOLUTE", token.KEYWORD_ABSOLUTE_TOKEN, "ABSOLUTE"},
		{"ACTION", token.KEYWORD_ACTION_TOKEN, "ACTION"},
		{"ADD", token.KEYWORD_ADD_TOKEN, "ADD"},
		{"ALL", token.KEYWORD_ALL_TOKEN, "ALL"},
		{"ALLOCATE", token.KEYWORD_ALLOCATE_TOKEN, "ALLOCATE"},
		{"ALTER", token.KEYWORD_ALTER_TOKEN, "ALTER"},
		{"AND", token.KEYWORD_AND_TOKEN, "AND"},
		{"ANY", token.KEYWORD_ANY_TOKEN, "ANY"},
		{"ARE", token.KEYWORD_ARE_TOKEN, "ARE"},
		{"AS", token.KEYWORD_AS_TOKEN, "AS"},
		{"ASC", token.KEYWORD_ASC_TOKEN, "ASC"},
		{"ASSERTION", token.KEYWORD_ASSERTION_TOKEN, "ASSERTION"},
		{"AT", token.KEYWORD_AT_TOKEN, "AT"},
		{"AUTHORIZATION", token.KEYWORD_AUTHORIZATION_TOKEN, "AUTHORIZATION"},
		{"AVG", token.KEYWORD_AVG_TOKEN, "AVG"},
		{"BEGIN", token.KEYWORD_BEGIN_TOKEN, "BEGIN"},
		{"BETWEEN", token.KEYWORD_BETWEEN_TOKEN, "BETWEEN"},
		{"BIT", token.KEYWORD_BIT_TOKEN, "BIT"},
		{"BIT_LENGTH", token.KEYWORD_BIT_LENGTH_TOKEN, "BIT_LENGTH"},
		{"BOTH", token.KEYWORD_BOTH_TOKEN, "BOTH"},
		{"BY", token.KEYWORD_BY_TOKEN, "BY"},
		{"CASCADE", token.KEYWORD_CASCADE_TOKEN, "CASCADE"},
		{"CASCADED", token.KEYWORD_CASCADED_TOKEN, "CASCADED"},
		{"CASE", token.KEYWORD_CASE_TOKEN, "CASE"},
		{"CAST", token.KEYWORD_CAST_TOKEN, "CAST"},
		{"CATALOG", token.KEYWORD_CATALOG_TOKEN, "CATALOG"},
		{"CHAR", token.KEYWORD_CHAR_TOKEN, "CHAR"},
		{"CHARACTER", token.KEYWORD_CHARACTER_TOKEN, "CHARACTER"},
		{"CHAR_LENGTH", token.KEYWORD_CHAR_LENGTH_TOKEN, "CHAR_LENGTH"},
		{"CHARACTER_LENGTH", token.KEYWORD_CHARACTER_LENGTH_TOKEN, "CHARACTER_LENGTH"},
		{"CHECK", token.KEYWORD_CHECK_TOKEN, "CHECK"},
		{"CLOSE", token.KEYWORD_CLOSE_TOKEN, "CLOSE"},
		{"COALESCE", token.KEYWORD_COALESCE_TOKEN, "COALESCE"},
		{"COLLATE", token.KEYWORD_COLLATE_TOKEN, "COLLATE"},
		{"COLLATION", token.KEYWORD_COLLATION_TOKEN, "COLLATION"},
		{"COLUMN", token.KEYWORD_COLUMN_TOKEN, "COLUMN"},
		{"COMMIT", token.KEYWORD_COMMIT_TOKEN, "COMMIT"},
		{"CONNECT", token.KEYWORD_CONNECT_TOKEN, "CONNECT"},
		{"CONNECTION", token.KEYWORD_CONNECTION_TOKEN, "CONNECTION"},
		{"CONSTRAINT", token.KEYWORD_CONSTRAINT_TOKEN, "CONSTRAINT"},
		{"CONSTRAINTS", token.KEYWORD_CONSTRAINTS_TOKEN, "CONSTRAINTS"},
		{"CONTINUE", token.KEYWORD_CONTINUE_TOKEN, "CONTINUE"},
		{"CONVERT", token.KEYWORD_CONVERT_TOKEN, "CONVERT"},
		{"CORRESPONDING", token.KEYWORD_CORRESPONDING_TOKEN, "CORRESPONDING"},
		{"COUNT", token.KEYWORD_COUNT_TOKEN, "COUNT"},
		{"CREATE", token.KEYWORD_CREATE_TOKEN, "CREATE"},
		{"CROSS", token.KEYWORD_CROSS_TOKEN, "CROSS"},
		{"CURRENT", token.KEYWORD_CURRENT_TOKEN, "CURRENT"},
		{"CURRENT_DATE", token.KEYWORD_CURRENT_DATE_TOKEN, "CURRENT_DATE"},
		{"CURRENT_TIME", token.KEYWORD_CURRENT_TIME_TOKEN, "CURRENT_TIME"},
		{"CURRENT_TIMESTAMP", token.KEYWORD_CURRENT_TIMESTAMP_TOKEN, "CURRENT_TIMESTAMP"},
		{"CURRENT_USER", token.KEYWORD_CURRENT_USER_TOKEN, "CURRENT_USER"},
		{"CURSOR", token.KEYWORD_CURSOR_TOKEN, "CURSOR"},
		{"DATE", token.KEYWORD_DATE_TOKEN, "DATE"},
		{"DAY", token.KEYWORD_DAY_TOKEN, "DAY"},
		{"DEALLOCATE", token.KEYWORD_DEALLOCATE_TOKEN, "DEALLOCATE"},
		{"DEC", token.KEYWORD_DEC_TOKEN, "DEC"},
		{"DECIMAL", token.KEYWORD_DECIMAL_TOKEN, "DECIMAL"},
		{"DECLARE", token.KEYWORD_DECLARE_TOKEN, "DECLARE"},
		{"DEFAULT", token.KEYWORD_DEFAULT_TOKEN, "DEFAULT"},
		{"DEFERRABLE", token.KEYWORD_DEFERRABLE_TOKEN, "DEFERRABLE"},
		{"DEFERRED", token.KEYWORD_DEFERRED_TOKEN, "DEFERRED"},
		{"DELETE", token.KEYWORD_DELETE_TOKEN, "DELETE"},
		{"DESC", token.KEYWORD_DESC_TOKEN, "DESC"},
		{"DESCRIBE", token.KEYWORD_DESCRIBE_TOKEN, "DESCRIBE"},
		{"DESCRIPTOR", token.KEYWORD_DESCRIPTOR_TOKEN, "DESCRIPTOR"},
		{"DIAGNOSTICS", token.KEYWORD_DIAGNOSTICS_TOKEN, "DIAGNOSTICS"},
		{"DISCONNECT", token.KEYWORD_DISCONNECT_TOKEN, "DISCONNECT"},
		{"DISTINCT", token.KEYWORD_DISTINCT_TOKEN, "DISTINCT"},
		{"DOMAIN", token.KEYWORD_DOMAIN_TOKEN, "DOMAIN"},
		{"DOUBLE", token.KEYWORD_DOUBLE_TOKEN, "DOUBLE"},
		{"DROP", token.KEYWORD_DROP_TOKEN, "DROP"},
		{"ELSE", token.KEYWORD_ELSE_TOKEN, "ELSE"},
		{"END", token.KEYWORD_END_TOKEN, "END"},
		{"END-EXEC", token.KEYWORD_END_EXEC_TOKEN, "END-EXEC"},
		{"ESCAPE", token.KEYWORD_ESCAPE_TOKEN, "ESCAPE"},
		{"EXCEPT", token.KEYWORD_EXCEPT_TOKEN, "EXCEPT"},
		{"EXCEPTION", token.KEYWORD_EXCEPTION_TOKEN, "EXCEPTION"},
		{"EXEC", token.KEYWORD_EXEC_TOKEN, "EXEC"},
		{"EXECUTE", token.KEYWORD_EXECUTE_TOKEN, "EXECUTE"},
		{"EXISTS", token.KEYWORD_EXISTS_TOKEN, "EXISTS"},
		{"EXTERNAL", token.KEYWORD_EXTERNAL_TOKEN, "EXTERNAL"},
		{"EXTRACT", token.KEYWORD_EXTRACT_TOKEN, "EXTRACT"},
		{"FALSE", token.KEYWORD_FALSE_TOKEN, "FALSE"},
		{"FETCH", token.KEYWORD_FETCH_TOKEN, "FETCH"},
		{"FIRST", token.KEYWORD_FIRST_TOKEN, "FIRST"},
		{"FLOAT", token.KEYWORD_FLOAT_TOKEN, "FLOAT"},
		{"FOR", token.KEYWORD_FOR_TOKEN, "FOR"},
		{"FOREIGN", token.KEYWORD_FOREIGN_TOKEN, "FOREIGN"},
		{"FOUND", token.KEYWORD_FOUND_TOKEN, "FOUND"},
		{"FROM", token.KEYWORD_FROM_TOKEN, "FROM"},
		{"FULL", token.KEYWORD_FULL_TOKEN, "FULL"},
		{"GET", token.KEYWORD_GET_TOKEN, "GET"},
		{"GLOBAL", token.KEYWORD_GLOBAL_TOKEN, "GLOBAL"},
		{"GO", token.KEYWORD_GO_TOKEN, "GO"},
		{"GOTO", token.KEYWORD_GOTO_TOKEN, "GOTO"},
		{"GRANT", token.KEYWORD_GRANT_TOKEN, "GRANT"},
		{"GROUP", token.KEYWORD_GROUP_TOKEN, "GROUP"},
		{"HAVING", token.KEYWORD_HAVING_TOKEN, "HAVING"},
		{"HOUR", token.KEYWORD_HOUR_TOKEN, "HOUR"},
		{"IDENTITY", token.KEYWORD_IDENTITY_TOKEN, "IDENTITY"},
		{"IMMEDIATE", token.KEYWORD_IMMEDIATE_TOKEN, "IMMEDIATE"},
		{"IN", token.KEYWORD_IN_TOKEN, "IN"},
		{"INDICATOR", token.KEYWORD_INDICATOR_TOKEN, "INDICATOR"},
		{"INITIALLY", token.KEYWORD_INITIALLY_TOKEN, "INITIALLY"},
		{"INNER", token.KEYWORD_INNER_TOKEN, "INNER"},
		{"INPUT", token.KEYWORD_INPUT_TOKEN, "INPUT"},
		{"INSENSITIVE", token.KEYWORD_INSENSITIVE_TOKEN, "INSENSITIVE"},
		{"INSERT", token.KEYWORD_INSERT_TOKEN, "INSERT"},
		{"INT", token.KEYWORD_INT_TOKEN, "INT"},
		{"INTEGER", token.KEYWORD_INTEGER_TOKEN, "INTEGER"},
		{"INTERSECT", token.KEYWORD_INTERSECT_TOKEN, "INTERSECT"},
		{"INTERVAL", token.KEYWORD_INTERVAL_TOKEN, "INTERVAL"},
		{"INTO", token.KEYWORD_INTO_TOKEN, "INTO"},
		{"IS", token.KEYWORD_IS_TOKEN, "IS"},
		{"ISOLATION", token.KEYWORD_ISOLATION_TOKEN, "ISOLATION"},
		{"JOIN", token.KEYWORD_JOIN_TOKEN, "JOIN"},
		{"KEY", token.KEYWORD_KEY_TOKEN, "KEY"},
		{"LANGUAGE", token.KEYWORD_LANGUAGE_TOKEN, "LANGUAGE"},
		{"LAST", token.KEYWORD_LAST_TOKEN, "LAST"},
		{"LEADING", token.KEYWORD_LEADING_TOKEN, "LEADING"},
		{"LEFT", token.KEYWORD_LEFT_TOKEN, "LEFT"},
		{"LEVEL", token.KEYWORD_LEVEL_TOKEN, "LEVEL"},
		{"LIKE", token.KEYWORD_LIKE_TOKEN, "LIKE"},
		{"LOCAL", token.KEYWORD_LOCAL_TOKEN, "LOCAL"},
		{"LOWER", token.KEYWORD_LOWER_TOKEN, "LOWER"},
		{"MATCH", token.KEYWORD_MATCH_TOKEN, "MATCH"},
		{"MAX", token.KEYWORD_MAX_TOKEN, "MAX"},
		{"MIN", token.KEYWORD_MIN_TOKEN, "MIN"},
		{"MINUTE", token.KEYWORD_MINUTE_TOKEN, "MINUTE"},
		{"MODULE", token.KEYWORD_MODULE_TOKEN, "MODULE"},
		{"MONTH", token.KEYWORD_MONTH_TOKEN, "MONTH"},
		{"NAMES", token.KEYWORD_NAMES_TOKEN, "NAMES"},
		{"NATIONAL", token.KEYWORD_NATIONAL_TOKEN, "NATIONAL"},
		{"NATURAL", token.KEYWORD_NATURAL_TOKEN, "NATURAL"},
		{"NCHAR", token.KEYWORD_NCHAR_TOKEN, "NCHAR"},
		{"NEXT", token.KEYWORD_NEXT_TOKEN, "NEXT"},
		{"NO", token.KEYWORD_NO_TOKEN, "NO"},
		{"NOT", token.KEYWORD_NOT_TOKEN, "NOT"},
		{"NULL", token.KEYWORD_NULL_TOKEN, "NULL"},
		{"NULLIF", token.KEYWORD_NULLIF_TOKEN, "NULLIF"},
		{"NUMERIC", token.KEYWORD_NUMERIC_TOKEN, "NUMERIC"},
		{"OCTET_LENGTH", token.KEYWORD_OCTET_LENGTH_TOKEN, "OCTET_LENGTH"},
		{"OF", token.KEYWORD_OF_TOKEN, "OF"},
		{"ON", token.KEYWORD_ON_TOKEN, "ON"},
		{"ONLY", token.KEYWORD_ONLY_TOKEN, "ONLY"},
		{"OPEN", token.KEYWORD_OPEN_TOKEN, "OPEN"},
		{"OPTION", token.KEYWORD_OPTION_TOKEN, "OPTION"},
		{"OR", token.KEYWORD_OR_TOKEN, "OR"},
		{"ORDER", token.KEYWORD_ORDER_TOKEN, "ORDER"},
		{"OUTER", token.KEYWORD_OUTER_TOKEN, "OUTER"},
		{"OUTPUT", token.KEYWORD_OUTPUT_TOKEN, "OUTPUT"},
		{"OVERLAPS", token.KEYWORD_OVERLAPS_TOKEN, "OVERLAPS"},
		{"PAD", token.KEYWORD_PAD_TOKEN, "PAD"},
		{"PARTIAL", token.KEYWORD_PARTIAL_TOKEN, "PARTIAL"},
		{"POSITION", token.KEYWORD_POSITION_TOKEN, "POSITION"},
		{"PRECISION", token.KEYWORD_PRECISION_TOKEN, "PRECISION"},
		{"PREPARE", token.KEYWORD_PREPARE_TOKEN, "PREPARE"},
		{"PRESERVE", token.KEYWORD_PRESERVE_TOKEN, "PRESERVE"},
		{"PRIMARY", token.KEYWORD_PRIMARY_TOKEN, "PRIMARY"},
		{"PRIOR", token.KEYWORD_PRIOR_TOKEN, "PRIOR"},
		{"PRIVILEGES", token.KEYWORD_PRIVILEGES_TOKEN, "PRIVILEGES"},
		{"PROCEDURE", token.KEYWORD_PROCEDURE_TOKEN, "PROCEDURE"},
		{"PUBLIC", token.KEYWORD_PUBLIC_TOKEN, "PUBLIC"},
		{"READ", token.KEYWORD_READ_TOKEN, "READ"},
		{"REAL", token.KEYWORD_REAL_TOKEN, "REAL"},
		{"REFERENCES", token.KEYWORD_REFERENCES_TOKEN, "REFERENCES"},
		{"RELATIVE", token.KEYWORD_RELATIVE_TOKEN, "RELATIVE"},
		{"RESTRICT", token.KEYWORD_RESTRICT_TOKEN, "RESTRICT"},
		{"REVOKE", token.KEYWORD_REVOKE_TOKEN, "REVOKE"},
		{"RIGHT", token.KEYWORD_RIGHT_TOKEN, "RIGHT"},
		{"ROLLBACK", token.KEYWORD_ROLLBACK_TOKEN, "ROLLBACK"},
		{"ROWS", token.KEYWORD_ROWS_TOKEN, "ROWS"},
		{"SCHEMA", token.KEYWORD_SCHEMA_TOKEN, "SCHEMA"},
		{"SCROLL", token.KEYWORD_SCROLL_TOKEN, "SCROLL"},
		{"SECOND", token.KEYWORD_SECOND_TOKEN, "SECOND"},
		{"SECTION", token.KEYWORD_SECTION_TOKEN, "SECTION"},
		{"SELECT", token.KEYWORD_SELECT_TOKEN, "SELECT"},
		{"SESSION", token.KEYWORD_SESSION_TOKEN, "SESSION"},
		{"SESSION_USER", token.KEYWORD_SESSION_USER_TOKEN, "SESSION_USER"},
		{"SET", token.KEYWORD_SET_TOKEN, "SET"},
		{"SIZE", token.KEYWORD_SIZE_TOKEN, "SIZE"},
		{"SMALLINT", token.KEYWORD_SMALLINT_TOKEN, "SMALLINT"},
		{"SOME", token.KEYWORD_SOME_TOKEN, "SOME"},
		{"SPACE", token.KEYWORD_SPACE_TOKEN, "SPACE"},
		{"SQL", token.KEYWORD_SQL_TOKEN, "SQL"},
		{"SQLCODE", token.KEYWORD_SQLCODE_TOKEN, "SQLCODE"},
		{"SQLERROR", token.KEYWORD_SQLERROR_TOKEN, "SQLERROR"},
		{"SQLSTATE", token.KEYWORD_SQLSTATE_TOKEN, "SQLSTATE"},
		{"SUBSTRING", token.KEYWORD_SUBSTRING_TOKEN, "SUBSTRING"},
		{"SUM", token.KEYWORD_SUM_TOKEN, "SUM"},
		{"SYSTEM_USER", token.KEYWORD_SYSTEM_USER_TOKEN, "SYSTEM_USER"},
		{"TABLE", token.KEYWORD_TABLE_TOKEN, "TABLE"},
		{"TEMPORARY", token.KEYWORD_TEMPORARY_TOKEN, "TEMPORARY"},
		{"THEN", token.KEYWORD_THEN_TOKEN, "THEN"},
		{"TIME", token.KEYWORD_TIME_TOKEN, "TIME"},
		{"TIMESTAMP", token.KEYWORD_TIMESTAMP_TOKEN, "TIMESTAMP"},
		{"TIMEZONE_HOUR", token.KEYWORD_TIMEZONE_HOUR_TOKEN, "TIMEZONE_HOUR"},
		{"TIMEZONE_MINUTE", token.KEYWORD_TIMEZONE_MINUTE_TOKEN, "TIMEZONE_MINUTE"},
		{"TO", token.KEYWORD_TO_TOKEN, "TO"},
		{"TRAILING", token.KEYWORD_TRAILING_TOKEN, "TRAILING"},
		{"TRANSACTION", token.KEYWORD_TRANSACTION_TOKEN, "TRANSACTION"},
		{"TRANSLATE", token.KEYWORD_TRANSLATE_TOKEN, "TRANSLATE"},
		{"TRANSLATION", token.KEYWORD_TRANSLATION_TOKEN, "TRANSLATION"},
		{"TRIM", token.KEYWORD_TRIM_TOKEN, "TRIM"},
		{"TRUE", token.KEYWORD_TRUE_TOKEN, "TRUE"},
		{"UNION", token.KEYWORD_UNION_TOKEN, "UNION"},
		{"UNIQUE", token.KEYWORD_UNIQUE_TOKEN, "UNIQUE"},
		{"UNKNOWN", token.KEYWORD_UNKNOWN_TOKEN, "UNKNOWN"},
		{"UPDATE", token.KEYWORD_UPDATE_TOKEN, "UPDATE"},
		{"UPPER", token.KEYWORD_UPPER_TOKEN, "UPPER"},
		{"USAGE", token.KEYWORD_USAGE_TOKEN, "USAGE"},
		{"USER", token.KEYWORD_USER_TOKEN, "USER"},
		{"USING", token.KEYWORD_USING_TOKEN, "USING"},
		{"VALUE", token.KEYWORD_VALUE_TOKEN, "VALUE"},
		{"VALUES", token.KEYWORD_VALUES_TOKEN, "VALUES"},
		{"VARCHAR", token.KEYWORD_VARCHAR_TOKEN, "VARCHAR"},
		{"VARYING", token.KEYWORD_VARYING_TOKEN, "VARYING"},
		{"VIEW", token.KEYWORD_VIEW_TOKEN, "VIEW"},
		{"WHEN", token.KEYWORD_WHEN_TOKEN, "WHEN"},
		{"WHENEVER", token.KEYWORD_WHENEVER_TOKEN, "WHENEVER"},
		{"WHERE", token.KEYWORD_WHERE_TOKEN, "WHERE"},
		{"WITH", token.KEYWORD_WITH_TOKEN, "WITH"},
		{"WORK", token.KEYWORD_WORK_TOKEN, "WORK"},
		{"WRITE", token.KEYWORD_WRITE_TOKEN, "WRITE"},
		{"YEAR", token.KEYWORD_YEAR_TOKEN, "YEAR"},
		{"ZONE", token.KEYWORD_ZONE_TOKEN, "ZONE"},
		{"ADA", token.KEYWORD_ADA_TOKEN, "ADA"},
		{"C", token.KEYWORD_C_TOKEN, "C"},
		{"CATALOG_NAME", token.KEYWORD_CATALOG_NAME_TOKEN, "CATALOG_NAME"},
		{"CHARACTER_SET_CATALOG", token.KEYWORD_CHARACTER_SET_CATALOG_TOKEN, "CHARACTER_SET_CATALOG"},
		{"CHARACTER_SET_NAME", token.KEYWORD_CHARACTER_SET_NAME_TOKEN, "CHARACTER_SET_NAME"},
		{"CHARACTER_SET_SCHEMA", token.KEYWORD_CHARACTER_SET_SCHEMA_TOKEN, "CHARACTER_SET_SCHEMA"},
		{"CLASS_ORIGIN", token.KEYWORD_CLASS_ORIGIN_TOKEN, "CLASS_ORIGIN"},
		{"COBOL", token.KEYWORD_COBOL_TOKEN, "COBOL"},
		{"COLLATION_CATALOG", token.KEYWORD_COLLATION_CATALOG_TOKEN, "COLLATION_CATALOG"},
		{"COLLATION_NAME", token.KEYWORD_COLLATION_NAME_TOKEN, "COLLATION_NAME"},
		{"COLLATION_SCHEMA", token.KEYWORD_COLLATION_SCHEMA_TOKEN, "COLLATION_SCHEMA"},
		{"COLUMN_NAME", token.KEYWORD_COLUMN_NAME_TOKEN, "COLUMN_NAME"},
		{"COMMAND_FUNCTION", token.KEYWORD_COMMAND_FUNCTION_TOKEN, "COMMAND_FUNCTION"},
		{"COMMITTED", token.KEYWORD_COMMITTED_TOKEN, "COMMITTED"},
		{"CONDITION_NUMBER", token.KEYWORD_CONDITION_NUMBER_TOKEN, "CONDITION_NUMBER"},
		{"CONNECTION_NAME", token.KEYWORD_CONNECTION_NAME_TOKEN, "CONNECTION_NAME"},
		{"CONSTRAINT_CATALOG", token.KEYWORD_CONSTRAINT_CATALOG_TOKEN, "CONSTRAINT_CATALOG"},
		{"CONSTRAINT_NAME", token.KEYWORD_CONSTRAINT_NAME_TOKEN, "CONSTRAINT_NAME"},
		{"CONSTRAINT_SCHEMA", token.KEYWORD_CONSTRAINT_SCHEMA_TOKEN, "CONSTRAINT_SCHEMA"},
		{"CURSOR_NAME", token.KEYWORD_CURSOR_NAME_TOKEN, "CURSOR_NAME"},
		{"DATA", token.KEYWORD_DATA_TOKEN, "DATA"},
		{"DATETIME_INTERVAL_CODE", token.KEYWORD_DATETIME_INTERVAL_CODE_TOKEN, "DATETIME_INTERVAL_CODE"},
		{"DATETIME_INTERVAL_PRECISION", token.KEYWORD_DATETIME_INTERVAL_PRECISION_TOKEN, "DATETIME_INTERVAL_PRECISION"},
		{"DYNAMIC_FUNCTION", token.KEYWORD_DYNAMIC_FUNCTION_TOKEN, "DYNAMIC_FUNCTION"},
		{"FORTRAN", token.KEYWORD_FORTRAN_TOKEN, "FORTRAN"},
		{"LENGTH", token.KEYWORD_LENGTH_TOKEN, "LENGTH"},
		{"MESSAGE_LENGTH", token.KEYWORD_MESSAGE_LENGTH_TOKEN, "MESSAGE_LENGTH"},
		{"MESSAGE_OCTET_LENGTH", token.KEYWORD_MESSAGE_OCTET_LENGTH_TOKEN, "MESSAGE_OCTET_LENGTH"},
		{"MESSAGE_TEXT", token.KEYWORD_MESSAGE_TEXT_TOKEN, "MESSAGE_TEXT"},
		{"MORE", token.KEYWORD_MORE_TOKEN, "MORE"},
		{"MUMPS", token.KEYWORD_MUMPS_TOKEN, "MUMPS"},
		{"NAME", token.KEYWORD_NAME_TOKEN, "NAME"},
		{"NULLABLE", token.KEYWORD_NULLABLE_TOKEN, "NULLABLE"},
		{"NUMBER", token.KEYWORD_NUMBER_TOKEN, "NUMBER"},
		{"PASCAL", token.KEYWORD_PASCAL_TOKEN, "PASCAL"},
		{"PLI", token.KEYWORD_PLI_TOKEN, "PLI"},
		{"REPEATABLE", token.KEYWORD_REPEATABLE_TOKEN, "REPEATABLE"},
		{"RETURNED_LENGTH", token.KEYWORD_RETURNED_LENGTH_TOKEN, "RETURNED_LENGTH"},
		{"RETURNED_OCTET_LENGTH", token.KEYWORD_RETURNED_OCTET_LENGTH_TOKEN, "RETURNED_OCTET_LENGTH"},
		{"RETURNED_SQLSTATE", token.KEYWORD_RETURNED_SQLSTATE_TOKEN, "RETURNED_SQLSTATE"},
		{"ROW_COUNT", token.KEYWORD_ROW_COUNT_TOKEN, "ROW_COUNT"},
		{"SCALE", token.KEYWORD_SCALE_TOKEN, "SCALE"},
		{"SCHEMA_NAME", token.KEYWORD_SCHEMA_NAME_TOKEN, "SCHEMA_NAME"},
		{"SERIALIZABLE", token.KEYWORD_SERIALIZABLE_TOKEN, "SERIALIZABLE"},
		{"SERVER_NAME", token.KEYWORD_SERVER_NAME_TOKEN, "SERVER_NAME"},
		{"SUBCLASS_ORIGIN", token.KEYWORD_SUBCLASS_ORIGIN_TOKEN, "SUBCLASS_ORIGIN"},
		{"TABLE_NAME", token.KEYWORD_TABLE_NAME_TOKEN, "TABLE_NAME"},
		{"TYPE", token.KEYWORD_TYPE_TOKEN, "TYPE"},
		{"UNCOMMITTED", token.KEYWORD_UNCOMMITTED_TOKEN, "UNCOMMITTED"},
		{"UNNAMED", token.KEYWORD_UNNAMED_TOKEN, "UNNAMED"},
		{"TBL", token.IDENTIFIER_TOKEN, "TBL"},
		{"age", token.IDENTIFIER_TOKEN, "age"},
		{"_sjis\"test\"", token.IDENTIFIER_TOKEN, "_sjis\"test\""},
	}

	for i, tt := range tests {
		l := New(tt.input)
		tok, li, err := l.readIdentifier()
		if err != nil {
			t.Errorf("[%d] Err error is not nil: %s (%s)\n", i, err, tt.input)
		} else {
			if tok != tt.expectedToken {
				t.Errorf("[%d] Err: Token type mistmatch (expected: %s, but got %s)\n", i, tt.expectedToken, tok)
			}
			if li != tt.expectedLiteral {
				t.Errorf("[%d] Err: Token Literal mistmatch (expected: %s, but got: %s)\n", i, tt.expectedLiteral, li)
			}
		}
	}
}

func TestNextToken(t *testing.T) {
	tests := []struct {
		input         string
		expectedToken token.TokenType
		isError       bool
	}{
		{"\"", token.DOUBLE_QUOTE_TOKEN, false},
		{"%", token.PERCENT_TOKEN, false},
		{"&", token.AMPERSAND_TOKEN, false},
		{"'", token.QUOTE_TOKEN, false},
		{"(", token.LEFT_PAREN_TOKEN, false},
		{")", token.RIGHT_PAREN_TOKEN, false},
		{"*", token.ASTERISK_TOKEN, false},
		{"+", token.PLUS_SIGN_TOKEN, false},
		{",", token.COMMA_TOKEN, false},
		{"-", token.MINUS_SIGN_TOKEN, false},
		{".", token.PERIOD_TOKEN, false},
		{"/", token.SOLIDAS_TOKEN, false},
		{"<>", token.NOT_EQUALS_OPERATOR_TOKEN, false},
		{">=", token.GREATER_THAN_OR_EQUALS_OPERATOR_TOKEN, false},
		{"<=", token.LESS_THAN_OR_EQUALS_OPERATOR_TOKEN, false},
		{"||", token.CONCATENATION_OPERATOR_TOKEN, false},
		{"..", token.DOUBLE_PERIOD_TOKEN, false},
		{"ABSOLUTE", token.KEYWORD_ABSOLUTE_TOKEN, false},
		{"ACTION", token.KEYWORD_ACTION_TOKEN, false},
		{"ADD", token.KEYWORD_ADD_TOKEN, false},
		{"ALL", token.KEYWORD_ALL_TOKEN, false},
		{"ALLOCATE", token.KEYWORD_ALLOCATE_TOKEN, false},
		{"ALTER", token.KEYWORD_ALTER_TOKEN, false},
		{"AND", token.KEYWORD_AND_TOKEN, false},
		{"ANY", token.KEYWORD_ANY_TOKEN, false},
		{"ARE", token.KEYWORD_ARE_TOKEN, false},
		{"AS", token.KEYWORD_AS_TOKEN, false},
		{"ASC", token.KEYWORD_ASC_TOKEN, false},
		{"ASSERTION", token.KEYWORD_ASSERTION_TOKEN, false},
		{"AT", token.KEYWORD_AT_TOKEN, false},
		{"AUTHORIZATION", token.KEYWORD_AUTHORIZATION_TOKEN, false},
		{"AVG", token.KEYWORD_AVG_TOKEN, false},
		{"BEGIN", token.KEYWORD_BEGIN_TOKEN, false},
		{"BETWEEN", token.KEYWORD_BETWEEN_TOKEN, false},
		{"BIT", token.KEYWORD_BIT_TOKEN, false},
		{"BIT_LENGTH", token.KEYWORD_BIT_LENGTH_TOKEN, false},
		{"BOTH", token.KEYWORD_BOTH_TOKEN, false},
		{"BY", token.KEYWORD_BY_TOKEN, false},
		{"ROW_COUNT", token.KEYWORD_ROW_COUNT_TOKEN, false},
		{"SCALE", token.KEYWORD_SCALE_TOKEN, false},
		{"SCHEMA_NAME", token.KEYWORD_SCHEMA_NAME_TOKEN, false},
		{"SERIALIZABLE", token.KEYWORD_SERIALIZABLE_TOKEN, false},
		{"SERVER_NAME", token.KEYWORD_SERVER_NAME_TOKEN, false},
		{"SUBCLASS_ORIGIN", token.KEYWORD_SUBCLASS_ORIGIN_TOKEN, false},
		{"TABLE_NAME", token.KEYWORD_TABLE_NAME_TOKEN, false},
		{"TYPE", token.KEYWORD_TYPE_TOKEN, false},
		{"UNCOMMITTED", token.KEYWORD_UNCOMMITTED_TOKEN, false},
		{"UNNAMED", token.KEYWORD_UNNAMED_TOKEN, false},
		{"TBL", token.IDENTIFIER_TOKEN, false},
		{"age", token.IDENTIFIER_TOKEN, false},
		{"1", token.NUMBER_TOKEN, false},
		{"123", token.NUMBER_TOKEN, false},
		{"123.", token.NUMBER_TOKEN, false},
		{"123.4", token.NUMBER_TOKEN, false},
		{".1", token.NUMBER_TOKEN, false},
		{"1.1.1", token.NUMBER_TOKEN, false},
		{"1..2", token.NUMBER_TOKEN, false},
		{"1..", token.NUMBER_TOKEN, false},
		{"1E+1", token.NUMBER_TOKEN, false},
		{"1E-1", token.NUMBER_TOKEN, false},
		{"1E1", token.NUMBER_TOKEN, false},
		{"12E10", token.NUMBER_TOKEN, false},
		{"12.3E-2", token.NUMBER_TOKEN, false},
		{"1.E+2", token.NUMBER_TOKEN, false},
		{".2E+2", token.NUMBER_TOKEN, false},
		{"1E+", token.NUMBER_TOKEN, true},
		{"B'0101'", token.NUMBER_TOKEN, false},
		{"B'0101' \t '0101'", token.NUMBER_TOKEN, false},
		{"B'0101' -- Comment '0101'", token.NUMBER_TOKEN, false},
		{"B'0101' -- Comment\n '0101'", token.NUMBER_TOKEN, false},
		{"B'0201' -- Comment\n '0101'", token.NUMBER_TOKEN, true},
		{"X'0A01'", token.NUMBER_TOKEN, false},
		{"X'0a01' \t '0101'", token.NUMBER_TOKEN, false},
		{"X'0A01' -- Comment '0101'", token.NUMBER_TOKEN, false},
		{"X'0A01' -- Comment\n '0101'", token.NUMBER_TOKEN, false},
		{"X'0K01' -- Comment\n '0101'", token.NUMBER_TOKEN, true},
		{"B'0101' \t SELECT", token.NUMBER_TOKEN, false},
		{"X'0101' \t SELECT", token.NUMBER_TOKEN, false},
		{"--- Comment\n", token.COMMENT_TOKEN, false},
		{"--- Comment", token.COMMENT_TOKEN, false},
		{"-----1+2\n", token.COMMENT_TOKEN, false},
		{"-- 1---2\n", token.COMMENT_TOKEN, false},
		{"-- Expect: None\n", token.COMMENT_TOKEN, false},
	}
	for i, tt := range tests {
		l := New(tt.input)
		tkn, err := l.nextToken()
		if err != nil && tt.isError == false {
			t.Errorf("[%d] Err error is not nil: %s (%s)\n", i, err, tt.input)
		}
		if err == nil && tt.isError == true {
			t.Errorf("[%d] Err error is nil (%s)\n", i, tt.input)
		}
		if tkn.Type != tt.expectedToken && tt.isError == false {
			t.Errorf("[%d] Err: Token Literal mistmatch (expected: %s, but got: %s)\n", i, tt.expectedToken, tkn.Type)
		}
	}
}

func zip(a, b []token.Token) ([][2]token.Token, error) {
	if len(a) != len(b) {
		return nil, errors.New(fmt.Sprintf("Length mismatch required:%d, but %d got.", len(a), len(b)))
	}
	r := make([][2]token.Token, len(a), len(b))
	for i, e := range a {
		r[i] = [2]token.Token{e, b[i]}
	}
	return r, nil
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		input         string
		expectedToken []token.Token
	}{
		{"  SELECT\n n,\n c * .1,\n -1.2\nFROM table1\nWHERE y>1981 AND m<>11;",
			[]token.Token{
				token.Token{Type: token.KEYWORD_SELECT_TOKEN, Literal: "SELECT"},
				token.Token{Type: token.IDENTIFIER_TOKEN, Literal: "n"},
				token.Token{Type: token.COMMA_TOKEN, Literal: ","},
				token.Token{Type: token.KEYWORD_C_TOKEN, Literal: "c"},
				token.Token{Type: token.ASTERISK_TOKEN, Literal: "*"},
				token.Token{Type: token.NUMBER_TOKEN, Literal: ".1"},
				token.Token{Type: token.COMMA_TOKEN, Literal: ","},
				token.Token{Type: token.MINUS_SIGN_TOKEN, Literal: "-"},
				token.Token{Type: token.NUMBER_TOKEN, Literal: "1.2"},
				token.Token{Type: token.KEYWORD_FROM_TOKEN, Literal: "FROM"},
				token.Token{Type: token.IDENTIFIER_TOKEN, Literal: "table1"},
				token.Token{Type: token.KEYWORD_WHERE_TOKEN, Literal: "WHERE"},
				token.Token{Type: token.IDENTIFIER_TOKEN, Literal: "y"},
				token.Token{Type: token.GREATER_THAN_OPERATOR_TOKEN, Literal: ">"},
				token.Token{Type: token.NUMBER_TOKEN, Literal: "1981"},
				token.Token{Type: token.KEYWORD_AND_TOKEN, Literal: "AND"},
				token.Token{Type: token.IDENTIFIER_TOKEN, Literal: "m"},
				token.Token{Type: token.NOT_EQUALS_OPERATOR_TOKEN, Literal: "<>"},
				token.Token{Type: token.NUMBER_TOKEN, Literal: "11"},
				token.Token{Type: token.SEMICOLON_TOKEN, Literal: ";"},
				token.Token{Type: token.EOF_TOKEN, Literal: ""},
			},
		},
		{"  SELECT\n END-EXEC END-EX",
			[]token.Token{
				token.Token{Type: token.KEYWORD_SELECT_TOKEN, Literal: "SELECT"},
				token.Token{Type: token.KEYWORD_END_EXEC_TOKEN, Literal: "END-EXEC"},
				token.Token{Type: token.KEYWORD_END_TOKEN, Literal: "END"},
				token.Token{Type: token.MINUS_SIGN_TOKEN, Literal: "-"},
				token.Token{Type: token.IDENTIFIER_TOKEN, Literal: "EX"},
				token.Token{Type: token.EOF_TOKEN, Literal: ""},
			},
		},
	}

	for i, tt := range tests {
		l := New(tt.input)
		tkn, err := l.Tokenize()
		if err != nil {
			t.Errorf("Tokenize Error %s\n", err)
			continue
		}
		pair, err := zip(tt.expectedToken, tkn)
		if err != nil {
			t.Errorf("Result Number Mismatch %s\n", err)
			continue
		}
		for j, p := range pair {
			if p[0].Type != p[1].Type {
				t.Errorf("[%d-%d] Err expected: %s, but got: %s\n", i, j, p[0].Type, p[1].Type)
			}
		}

	}
}
