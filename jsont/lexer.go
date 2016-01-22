package jsont

import (
	"unicode"
)

// RW
func is_identifier_first(c rune) bool {
	return unicode.IsUpper(c) || unicode.IsLower(c) || c == '_'
}

func is_identifier(c rune) bool {
	return is_identifier_first(c) || unicode.IsDigit(c)
}

func is_symbol(c rune) bool {
	switch c {
	case '&', '|', '^', '=', '<', '>', '*', '/', '%', '#':
		return true
	}
	return false
}
