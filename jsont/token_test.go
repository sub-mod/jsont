package jsont_test

import (
	"strings"
	"testing"

	"github.com/sub-mod/jsont"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok jsont.Kind
		lit string
	}{
		{s: `.`, tok: jsont.COLON},
	}

	for i, tt := range tests {
		s := jsont.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)

		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)

		}

	}

}
