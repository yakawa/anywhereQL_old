package token

import (
	"testing"
)

func TestLookupKeyword(t *testing.T) {
	tests := []struct {
		input           string
		expectedType    TokenType
		expectedLiteral string
	}{
		{"ABSOLUTE", KEYWORD_ABSOLUTE_TOKEN, "ABSOLUTE"},
		{"ACTION", KEYWORD_ACTION_TOKEN, "ACTION"},
		{"ADD", KEYWORD_ADD_TOKEN, "ADD"},
		{"ALL", KEYWORD_ALL_TOKEN, "ALL"},
		{"ALLOCATE", KEYWORD_ALLOCATE_TOKEN, "ALLOCATE"},
		{"ALTER", KEYWORD_ALTER_TOKEN, "ALTER"},
		{"AND", KEYWORD_AND_TOKEN, "AND"},
		{"ANY", KEYWORD_ANY_TOKEN, "ANY"},
		{"ARE", KEYWORD_ARE_TOKEN, "ARE"},
		{"AS", KEYWORD_AS_TOKEN, "AS"},
		{"ASC", KEYWORD_ASC_TOKEN, "ASC"},
		{"ASSERTION", KEYWORD_ASSERTION_TOKEN, "ASSERTION"},
		{"AT", KEYWORD_AT_TOKEN, "AT"},
		{"AUTHORIZATION", KEYWORD_AUTHORIZATION_TOKEN, "AUTHORIZATION"},
		{"AVG", KEYWORD_AVG_TOKEN, "AVG"},
		{"BEGIN", KEYWORD_BEGIN_TOKEN, "BEGIN"},
		{"BETWEEN", KEYWORD_BETWEEN_TOKEN, "BETWEEN"},
		{"BIT", KEYWORD_BIT_TOKEN, "BIT"},
		{"BIT_LENGTH", KEYWORD_BIT_LENGTH_TOKEN, "BIT_LENGTH"},
		{"BOTH", KEYWORD_BOTH_TOKEN, "BOTH"},
		{"BY", KEYWORD_BY_TOKEN, "BY"},
		{"CASCADE", KEYWORD_CASCADE_TOKEN, "CASCADE"},
		{"CASCADED", KEYWORD_CASCADED_TOKEN, "CASCADED"},
		{"CASE", KEYWORD_CASE_TOKEN, "CASE"},
		{"CAST", KEYWORD_CAST_TOKEN, "CAST"},
		{"CATALOG", KEYWORD_CATALOG_TOKEN, "CATALOG"},
		{"CHAR", KEYWORD_CHAR_TOKEN, "CHAR"},
		{"CHARACTER", KEYWORD_CHARACTER_TOKEN, "CHARACTER"},
		{"CHAR_LENGTH", KEYWORD_CHAR_LENGTH_TOKEN, "CHAR_LENGTH"},
		{"CHARACTER_LENGTH", KEYWORD_CHARACTER_LENGTH_TOKEN, "CHARACTER_LENGTH"},
		{"CHECK", KEYWORD_CHECK_TOKEN, "CHECK"},
		{"CLOSE", KEYWORD_CLOSE_TOKEN, "CLOSE"},
		{"COALESCE", KEYWORD_COALESCE_TOKEN, "COALESCE"},
		{"COLLATE", KEYWORD_COLLATE_TOKEN, "COLLATE"},
		{"COLLATION", KEYWORD_COLLATION_TOKEN, "COLLATION"},
		{"COLUMN", KEYWORD_COLUMN_TOKEN, "COLUMN"},
		{"COMMIT", KEYWORD_COMMIT_TOKEN, "COMMIT"},
		{"CONNECT", KEYWORD_CONNECT_TOKEN, "CONNECT"},
		{"CONNECTION", KEYWORD_CONNECTION_TOKEN, "CONNECTION"},
		{"CONSTRAINT", KEYWORD_CONSTRAINT_TOKEN, "CONSTRAINT"},
		{"CONSTRAINTS", KEYWORD_CONSTRAINTS_TOKEN, "CONSTRAINTS"},
		{"CONTINUE", KEYWORD_CONTINUE_TOKEN, "CONTINUE"},
		{"CONVERT", KEYWORD_CONVERT_TOKEN, "CONVERT"},
		{"CORRESPONDING", KEYWORD_CORRESPONDING_TOKEN, "CORRESPONDING"},
		{"COUNT", KEYWORD_COUNT_TOKEN, "COUNT"},
		{"CREATE", KEYWORD_CREATE_TOKEN, "CREATE"},
		{"CROSS", KEYWORD_CROSS_TOKEN, "CROSS"},
		{"CURRENT", KEYWORD_CURRENT_TOKEN, "CURRENT"},
		{"CURRENT_DATE", KEYWORD_CURRENT_DATE_TOKEN, "CURRENT_DATE"},
		{"CURRENT_TIME", KEYWORD_CURRENT_TIME_TOKEN, "CURRENT_TIME"},
		{"CURRENT_TIMESTAMP", KEYWORD_CURRENT_TIMESTAMP_TOKEN, "CURRENT_TIMESTAMP"},
		{"CURRENT_USER", KEYWORD_CURRENT_USER_TOKEN, "CURRENT_USER"},
		{"CURSOR", KEYWORD_CURSOR_TOKEN, "CURSOR"},
		{"DATE", KEYWORD_DATE_TOKEN, "DATE"},
		{"DAY", KEYWORD_DAY_TOKEN, "DAY"},
		{"DEALLOCATE", KEYWORD_DEALLOCATE_TOKEN, "DEALLOCATE"},
		{"DEC", KEYWORD_DEC_TOKEN, "DEC"},
		{"DECIMAL", KEYWORD_DECIMAL_TOKEN, "DECIMAL"},
		{"DECLARE", KEYWORD_DECLARE_TOKEN, "DECLARE"},
		{"DEFAULT", KEYWORD_DEFAULT_TOKEN, "DEFAULT"},
		{"DEFERRABLE", KEYWORD_DEFERRABLE_TOKEN, "DEFERRABLE"},
		{"DEFERRED", KEYWORD_DEFERRED_TOKEN, "DEFERRED"},
		{"DELETE", KEYWORD_DELETE_TOKEN, "DELETE"},
		{"DESC", KEYWORD_DESC_TOKEN, "DESC"},
		{"DESCRIBE", KEYWORD_DESCRIBE_TOKEN, "DESCRIBE"},
		{"DESCRIPTOR", KEYWORD_DESCRIPTOR_TOKEN, "DESCRIPTOR"},
		{"DIAGNOSTICS", KEYWORD_DIAGNOSTICS_TOKEN, "DIAGNOSTICS"},
		{"DISCONNECT", KEYWORD_DISCONNECT_TOKEN, "DISCONNECT"},
		{"DISTINCT", KEYWORD_DISTINCT_TOKEN, "DISTINCT"},
		{"DOMAIN", KEYWORD_DOMAIN_TOKEN, "DOMAIN"},
		{"DOUBLE", KEYWORD_DOUBLE_TOKEN, "DOUBLE"},
		{"DROP", KEYWORD_DROP_TOKEN, "DROP"},
		{"ELSE", KEYWORD_ELSE_TOKEN, "ELSE"},
		{"END", KEYWORD_END_TOKEN, "END"},
		{"END-EXEC", KEYWORD_END_EXEC_TOKEN, "END-EXEC"},
		{"ESCAPE", KEYWORD_ESCAPE_TOKEN, "ESCAPE"},
		{"EXCEPT", KEYWORD_EXCEPT_TOKEN, "EXCEPT"},
		{"EXCEPTION", KEYWORD_EXCEPTION_TOKEN, "EXCEPTION"},
		{"EXEC", KEYWORD_EXEC_TOKEN, "EXEC"},
		{"EXECUTE", KEYWORD_EXECUTE_TOKEN, "EXECUTE"},
		{"EXISTS", KEYWORD_EXISTS_TOKEN, "EXISTS"},
		{"EXTERNAL", KEYWORD_EXTERNAL_TOKEN, "EXTERNAL"},
		{"EXTRACT", KEYWORD_EXTRACT_TOKEN, "EXTRACT"},
		{"FALSE", KEYWORD_FALSE_TOKEN, "FALSE"},
		{"FETCH", KEYWORD_FETCH_TOKEN, "FETCH"},
		{"FIRST", KEYWORD_FIRST_TOKEN, "FIRST"},
		{"FLOAT", KEYWORD_FLOAT_TOKEN, "FLOAT"},
		{"FOR", KEYWORD_FOR_TOKEN, "FOR"},
		{"FOREIGN", KEYWORD_FOREIGN_TOKEN, "FOREIGN"},
		{"FOUND", KEYWORD_FOUND_TOKEN, "FOUND"},
		{"FROM", KEYWORD_FROM_TOKEN, "FROM"},
		{"FULL", KEYWORD_FULL_TOKEN, "FULL"},
		{"GET", KEYWORD_GET_TOKEN, "GET"},
		{"GLOBAL", KEYWORD_GLOBAL_TOKEN, "GLOBAL"},
		{"GO", KEYWORD_GO_TOKEN, "GO"},
		{"GOTO", KEYWORD_GOTO_TOKEN, "GOTO"},
		{"GRANT", KEYWORD_GRANT_TOKEN, "GRANT"},
		{"GROUP", KEYWORD_GROUP_TOKEN, "GROUP"},
		{"HAVING", KEYWORD_HAVING_TOKEN, "HAVING"},
		{"HOUR", KEYWORD_HOUR_TOKEN, "HOUR"},
		{"IDENTITY", KEYWORD_IDENTITY_TOKEN, "IDENTITY"},
		{"IMMEDIATE", KEYWORD_IMMEDIATE_TOKEN, "IMMEDIATE"},
		{"IN", KEYWORD_IN_TOKEN, "IN"},
		{"INDICATOR", KEYWORD_INDICATOR_TOKEN, "INDICATOR"},
		{"INITIALLY", KEYWORD_INITIALLY_TOKEN, "INITIALLY"},
		{"INNER", KEYWORD_INNER_TOKEN, "INNER"},
		{"INPUT", KEYWORD_INPUT_TOKEN, "INPUT"},
		{"INSENSITIVE", KEYWORD_INSENSITIVE_TOKEN, "INSENSITIVE"},
		{"INSERT", KEYWORD_INSERT_TOKEN, "INSERT"},
		{"INT", KEYWORD_INT_TOKEN, "INT"},
		{"INTEGER", KEYWORD_INTEGER_TOKEN, "INTEGER"},
		{"INTERSECT", KEYWORD_INTERSECT_TOKEN, "INTERSECT"},
		{"INTERVAL", KEYWORD_INTERVAL_TOKEN, "INTERVAL"},
		{"INTO", KEYWORD_INTO_TOKEN, "INTO"},
		{"IS", KEYWORD_IS_TOKEN, "IS"},
		{"ISOLATION", KEYWORD_ISOLATION_TOKEN, "ISOLATION"},
		{"JOIN", KEYWORD_JOIN_TOKEN, "JOIN"},
		{"KEY", KEYWORD_KEY_TOKEN, "KEY"},
		{"LANGUAGE", KEYWORD_LANGUAGE_TOKEN, "LANGUAGE"},
		{"LAST", KEYWORD_LAST_TOKEN, "LAST"},
		{"LEADING", KEYWORD_LEADING_TOKEN, "LEADING"},
		{"LEFT", KEYWORD_LEFT_TOKEN, "LEFT"},
		{"LEVEL", KEYWORD_LEVEL_TOKEN, "LEVEL"},
		{"LIKE", KEYWORD_LIKE_TOKEN, "LIKE"},
		{"LOCAL", KEYWORD_LOCAL_TOKEN, "LOCAL"},
		{"LOWER", KEYWORD_LOWER_TOKEN, "LOWER"},
		{"MATCH", KEYWORD_MATCH_TOKEN, "MATCH"},
		{"MAX", KEYWORD_MAX_TOKEN, "MAX"},
		{"MIN", KEYWORD_MIN_TOKEN, "MIN"},
		{"MINUTE", KEYWORD_MINUTE_TOKEN, "MINUTE"},
		{"MODULE", KEYWORD_MODULE_TOKEN, "MODULE"},
		{"MONTH", KEYWORD_MONTH_TOKEN, "MONTH"},
		{"NAMES", KEYWORD_NAMES_TOKEN, "NAMES"},
		{"NATIONAL", KEYWORD_NATIONAL_TOKEN, "NATIONAL"},
		{"NATURAL", KEYWORD_NATURAL_TOKEN, "NATURAL"},
		{"NCHAR", KEYWORD_NCHAR_TOKEN, "NCHAR"},
		{"NEXT", KEYWORD_NEXT_TOKEN, "NEXT"},
		{"NO", KEYWORD_NO_TOKEN, "NO"},
		{"NOT", KEYWORD_NOT_TOKEN, "NOT"},
		{"NULL", KEYWORD_NULL_TOKEN, "NULL"},
		{"NULLIF", KEYWORD_NULLIF_TOKEN, "NULLIF"},
		{"NUMERIC", KEYWORD_NUMERIC_TOKEN, "NUMERIC"},
		{"OCTET_LENGTH", KEYWORD_OCTET_LENGTH_TOKEN, "OCTET_LENGTH"},
		{"OF", KEYWORD_OF_TOKEN, "OF"},
		{"ON", KEYWORD_ON_TOKEN, "ON"},
		{"ONLY", KEYWORD_ONLY_TOKEN, "ONLY"},
		{"OPEN", KEYWORD_OPEN_TOKEN, "OPEN"},
		{"OPTION", KEYWORD_OPTION_TOKEN, "OPTION"},
		{"OR", KEYWORD_OR_TOKEN, "OR"},
		{"ORDER", KEYWORD_ORDER_TOKEN, "ORDER"},
		{"OUTER", KEYWORD_OUTER_TOKEN, "OUTER"},
		{"OUTPUT", KEYWORD_OUTPUT_TOKEN, "OUTPUT"},
		{"OVERLAPS", KEYWORD_OVERLAPS_TOKEN, "OVERLAPS"},
		{"PAD", KEYWORD_PAD_TOKEN, "PAD"},
		{"PARTIAL", KEYWORD_PARTIAL_TOKEN, "PARTIAL"},
		{"POSITION", KEYWORD_POSITION_TOKEN, "POSITION"},
		{"PRECISION", KEYWORD_PRECISION_TOKEN, "PRECISION"},
		{"PREPARE", KEYWORD_PREPARE_TOKEN, "PREPARE"},
		{"PRESERVE", KEYWORD_PRESERVE_TOKEN, "PRESERVE"},
		{"PRIMARY", KEYWORD_PRIMARY_TOKEN, "PRIMARY"},
		{"PRIOR", KEYWORD_PRIOR_TOKEN, "PRIOR"},
		{"PRIVILEGES", KEYWORD_PRIVILEGES_TOKEN, "PRIVILEGES"},
		{"PROCEDURE", KEYWORD_PROCEDURE_TOKEN, "PROCEDURE"},
		{"PUBLIC", KEYWORD_PUBLIC_TOKEN, "PUBLIC"},
		{"READ", KEYWORD_READ_TOKEN, "READ"},
		{"REAL", KEYWORD_REAL_TOKEN, "REAL"},
		{"REFERENCES", KEYWORD_REFERENCES_TOKEN, "REFERENCES"},
		{"RELATIVE", KEYWORD_RELATIVE_TOKEN, "RELATIVE"},
		{"RESTRICT", KEYWORD_RESTRICT_TOKEN, "RESTRICT"},
		{"REVOKE", KEYWORD_REVOKE_TOKEN, "REVOKE"},
		{"RIGHT", KEYWORD_RIGHT_TOKEN, "RIGHT"},
		{"ROLLBACK", KEYWORD_ROLLBACK_TOKEN, "ROLLBACK"},
		{"ROWS", KEYWORD_ROWS_TOKEN, "ROWS"},
		{"SCHEMA", KEYWORD_SCHEMA_TOKEN, "SCHEMA"},
		{"SCROLL", KEYWORD_SCROLL_TOKEN, "SCROLL"},
		{"SECOND", KEYWORD_SECOND_TOKEN, "SECOND"},
		{"SECTION", KEYWORD_SECTION_TOKEN, "SECTION"},
		{"SELECT", KEYWORD_SELECT_TOKEN, "SELECT"},
		{"SESSION", KEYWORD_SESSION_TOKEN, "SESSION"},
		{"SESSION_USER", KEYWORD_SESSION_USER_TOKEN, "SESSION_USER"},
		{"SET", KEYWORD_SET_TOKEN, "SET"},
		{"SIZE", KEYWORD_SIZE_TOKEN, "SIZE"},
		{"SMALLINT", KEYWORD_SMALLINT_TOKEN, "SMALLINT"},
		{"SOME", KEYWORD_SOME_TOKEN, "SOME"},
		{"SPACE", KEYWORD_SPACE_TOKEN, "SPACE"},
		{"SQL", KEYWORD_SQL_TOKEN, "SQL"},
		{"SQLCODE", KEYWORD_SQLCODE_TOKEN, "SQLCODE"},
		{"SQLERROR", KEYWORD_SQLERROR_TOKEN, "SQLERROR"},
		{"SQLSTATE", KEYWORD_SQLSTATE_TOKEN, "SQLSTATE"},
		{"SUBSTRING", KEYWORD_SUBSTRING_TOKEN, "SUBSTRING"},
		{"SUM", KEYWORD_SUM_TOKEN, "SUM"},
		{"SYSTEM_USER", KEYWORD_SYSTEM_USER_TOKEN, "SYSTEM_USER"},
		{"TABLE", KEYWORD_TABLE_TOKEN, "TABLE"},
		{"TEMPORARY", KEYWORD_TEMPORARY_TOKEN, "TEMPORARY"},
		{"THEN", KEYWORD_THEN_TOKEN, "THEN"},
		{"TIME", KEYWORD_TIME_TOKEN, "TIME"},
		{"TIMESTAMP", KEYWORD_TIMESTAMP_TOKEN, "TIMESTAMP"},
		{"TIMEZONE_HOUR", KEYWORD_TIMEZONE_HOUR_TOKEN, "TIMEZONE_HOUR"},
		{"TIMEZONE_MINUTE", KEYWORD_TIMEZONE_MINUTE_TOKEN, "TIMEZONE_MINUTE"},
		{"TO", KEYWORD_TO_TOKEN, "TO"},
		{"TRAILING", KEYWORD_TRAILING_TOKEN, "TRAILING"},
		{"TRANSACTION", KEYWORD_TRANSACTION_TOKEN, "TRANSACTION"},
		{"TRANSLATE", KEYWORD_TRANSLATE_TOKEN, "TRANSLATE"},
		{"TRANSLATION", KEYWORD_TRANSLATION_TOKEN, "TRANSLATION"},
		{"TRIM", KEYWORD_TRIM_TOKEN, "TRIM"},
		{"TRUE", KEYWORD_TRUE_TOKEN, "TRUE"},
		{"UNION", KEYWORD_UNION_TOKEN, "UNION"},
		{"UNIQUE", KEYWORD_UNIQUE_TOKEN, "UNIQUE"},
		{"UNKNOWN", KEYWORD_UNKNOWN_TOKEN, "UNKNOWN"},
		{"UPDATE", KEYWORD_UPDATE_TOKEN, "UPDATE"},
		{"UPPER", KEYWORD_UPPER_TOKEN, "UPPER"},
		{"USAGE", KEYWORD_USAGE_TOKEN, "USAGE"},
		{"USER", KEYWORD_USER_TOKEN, "USER"},
		{"USING", KEYWORD_USING_TOKEN, "USING"},
		{"VALUE", KEYWORD_VALUE_TOKEN, "VALUE"},
		{"VALUES", KEYWORD_VALUES_TOKEN, "VALUES"},
		{"VARCHAR", KEYWORD_VARCHAR_TOKEN, "VARCHAR"},
		{"VARYING", KEYWORD_VARYING_TOKEN, "VARYING"},
		{"VIEW", KEYWORD_VIEW_TOKEN, "VIEW"},
		{"WHEN", KEYWORD_WHEN_TOKEN, "WHEN"},
		{"WHENEVER", KEYWORD_WHENEVER_TOKEN, "WHENEVER"},
		{"WHERE", KEYWORD_WHERE_TOKEN, "WHERE"},
		{"WITH", KEYWORD_WITH_TOKEN, "WITH"},
		{"WORK", KEYWORD_WORK_TOKEN, "WORK"},
		{"WRITE", KEYWORD_WRITE_TOKEN, "WRITE"},
		{"YEAR", KEYWORD_YEAR_TOKEN, "YEAR"},
		{"ZONE", KEYWORD_ZONE_TOKEN, "ZONE"},
		{"ADA", KEYWORD_ADA_TOKEN, "ADA"},
		{"C", KEYWORD_C_TOKEN, "C"},
		{"CATALOG_NAME", KEYWORD_CATALOG_NAME_TOKEN, "CATALOG_NAME"},
		{"CHARACTER_SET_CATALOG", KEYWORD_CHARACTER_SET_CATALOG_TOKEN, "CHARACTER_SET_CATALOG"},
		{"CHARACTER_SET_NAME", KEYWORD_CHARACTER_SET_NAME_TOKEN, "CHARACTER_SET_NAME"},
		{"CHARACTER_SET_SCHEMA", KEYWORD_CHARACTER_SET_SCHEMA_TOKEN, "CHARACTER_SET_SCHEMA"},
		{"CLASS_ORIGIN", KEYWORD_CLASS_ORIGIN_TOKEN, "CLASS_ORIGIN"},
		{"COBOL", KEYWORD_COBOL_TOKEN, "COBOL"},
		{"COLLATION_CATALOG", KEYWORD_COLLATION_CATALOG_TOKEN, "COLLATION_CATALOG"},
		{"COLLATION_NAME", KEYWORD_COLLATION_NAME_TOKEN, "COLLATION_NAME"},
		{"COLLATION_SCHEMA", KEYWORD_COLLATION_SCHEMA_TOKEN, "COLLATION_SCHEMA"},
		{"COLUMN_NAME", KEYWORD_COLUMN_NAME_TOKEN, "COLUMN_NAME"},
		{"COMMAND_FUNCTION", KEYWORD_COMMAND_FUNCTION_TOKEN, "COMMAND_FUNCTION"},
		{"COMMITTED", KEYWORD_COMMITTED_TOKEN, "COMMITTED"},
		{"CONDITION_NUMBER", KEYWORD_CONDITION_NUMBER_TOKEN, "CONDITION_NUMBER"},
		{"CONNECTION_NAME", KEYWORD_CONNECTION_NAME_TOKEN, "CONNECTION_NAME"},
		{"CONSTRAINT_CATALOG", KEYWORD_CONSTRAINT_CATALOG_TOKEN, "CONSTRAINT_CATALOG"},
		{"CONSTRAINT_NAME", KEYWORD_CONSTRAINT_NAME_TOKEN, "CONSTRAINT_NAME"},
		{"CONSTRAINT_SCHEMA", KEYWORD_CONSTRAINT_SCHEMA_TOKEN, "CONSTRAINT_SCHEMA"},
		{"CURSOR_NAME", KEYWORD_CURSOR_NAME_TOKEN, "CURSOR_NAME"},
		{"DATA", KEYWORD_DATA_TOKEN, "DATA"},
		{"DATETIME_INTERVAL_CODE", KEYWORD_DATETIME_INTERVAL_CODE_TOKEN, "DATETIME_INTERVAL_CODE"},
		{"DATETIME_INTERVAL_PRECISION", KEYWORD_DATETIME_INTERVAL_PRECISION_TOKEN, "DATETIME_INTERVAL_PRECISION"},
		{"DYNAMIC_FUNCTION", KEYWORD_DYNAMIC_FUNCTION_TOKEN, "DYNAMIC_FUNCTION"},
		{"FORTRAN", KEYWORD_FORTRAN_TOKEN, "FORTRAN"},
		{"LENGTH", KEYWORD_LENGTH_TOKEN, "LENGTH"},
		{"MESSAGE_LENGTH", KEYWORD_MESSAGE_LENGTH_TOKEN, "MESSAGE_LENGTH"},
		{"MESSAGE_OCTET_LENGTH", KEYWORD_MESSAGE_OCTET_LENGTH_TOKEN, "MESSAGE_OCTET_LENGTH"},
		{"MESSAGE_TEXT", KEYWORD_MESSAGE_TEXT_TOKEN, "MESSAGE_TEXT"},
		{"MORE", KEYWORD_MORE_TOKEN, "MORE"},
		{"MUMPS", KEYWORD_MUMPS_TOKEN, "MUMPS"},
		{"NAME", KEYWORD_NAME_TOKEN, "NAME"},
		{"NULLABLE", KEYWORD_NULLABLE_TOKEN, "NULLABLE"},
		{"NUMBER", KEYWORD_NUMBER_TOKEN, "NUMBER"},
		{"PASCAL", KEYWORD_PASCAL_TOKEN, "PASCAL"},
		{"PLI", KEYWORD_PLI_TOKEN, "PLI"},
		{"REPEATABLE", KEYWORD_REPEATABLE_TOKEN, "REPEATABLE"},
		{"RETURNED_LENGTH", KEYWORD_RETURNED_LENGTH_TOKEN, "RETURNED_LENGTH"},
		{"RETURNED_OCTET_LENGTH", KEYWORD_RETURNED_OCTET_LENGTH_TOKEN, "RETURNED_OCTET_LENGTH"},
		{"RETURNED_SQLSTATE", KEYWORD_RETURNED_SQLSTATE_TOKEN, "RETURNED_SQLSTATE"},
		{"ROW_COUNT", KEYWORD_ROW_COUNT_TOKEN, "ROW_COUNT"},
		{"SCALE", KEYWORD_SCALE_TOKEN, "SCALE"},
		{"SCHEMA_NAME", KEYWORD_SCHEMA_NAME_TOKEN, "SCHEMA_NAME"},
		{"SERIALIZABLE", KEYWORD_SERIALIZABLE_TOKEN, "SERIALIZABLE"},
		{"SERVER_NAME", KEYWORD_SERVER_NAME_TOKEN, "SERVER_NAME"},
		{"SUBCLASS_ORIGIN", KEYWORD_SUBCLASS_ORIGIN_TOKEN, "SUBCLASS_ORIGIN"},
		{"TABLE_NAME", KEYWORD_TABLE_NAME_TOKEN, "TABLE_NAME"},
		{"TYPE", KEYWORD_TYPE_TOKEN, "TYPE"},
		{"UNCOMMITTED", KEYWORD_UNCOMMITTED_TOKEN, "UNCOMMITTED"},
		{"UNNAMED", KEYWORD_UNNAMED_TOKEN, "UNNAMED"},
		{"TBL", IDENTIFIER_TOKEN, "TBL"},
		{"age", IDENTIFIER_TOKEN, "age"},
	}

	for i, tt := range tests {
		tok := LookupKeyword(tt.input)
		if tok != tt.expectedType {
			t.Errorf("[%d] Failed expect Type: %s ,but got %s\n", i, tt.expectedType, tok)
		}
	}
}

func TestDebug(t *testing.T) {
	tests := []struct {
		input Token
	}{
		{Token{Type: KEYWORD_SELECT_TOKEN, Literal: "SELECT"}},
	}

	for i, tt := range tests {
		r := tt.input.Debug()
		if r == "" {
			t.Errorf("[%d]Output is ignore\n", i)
		}
	}
}