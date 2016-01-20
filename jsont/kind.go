package jsont

type Kind int

const (
	BRACE_L = iota
	BRACE_R
	BRACKET_L
	BRACKET_R
	COLON
	COMMA
	DOLLAR
	EOF
)
