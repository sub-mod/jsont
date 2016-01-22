package jsont

type FodderElementKind int
type TokenKind int
type Fodder []FodderElement


const (
	WHITESPACE FodderElementKind = iota
	COMMENT_C
	COMMENT_CPP
	COMMENT_HASH
)

const (
	// Symbols
	BRACE_L TokenKind = iota
	BRACE_R
	BRACKET_L
	BRACKET_R
	COLON
	COMMA
	DOLLAR
	DOT
	PAREN_L
	PAREN_R
	SEMICOLON

	// Arbitrary length lexemes
	IDENTIFIER
	NUMBER
	OPERATOR
	STRING_DOUBLE
	STRING_SINGLE
	STRING_BLOCK

	// Keywords
	ASSERT
	ELSE
	ERROR
	FALSE
	FOR
	FUNCTION
	IF
	IMPORT
	IMPORTSTR
	IN
	LOCAL
	NULL_LIT
	TAILSTRICT
	THEN
	SELF
	SUPER
	TRUE
	// A special token that holds line/column information about the end of the file.
	END_OF_FILE
)

type FodderElement struct {
	kind FodderElementKind
	data string
}



type Token struct {

	fodder Fodder

    /** Content of the token if it wasn't a keyword. */
    data string

    /** If kind == STRING_BLOCK then stores the sequence of whitespace that indented the block. */
    stringBlockIndent string

    /** If kind == STRING_BLOCK then stores the sequence of whitespace that indented the end of
     * the block. 
     *
     * This is always fewer whitespace characters than in stringBlockIndent.
     */
    stringBlockTermIndent string

}
