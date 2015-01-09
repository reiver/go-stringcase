package stringcase


import "github.com/reiver/go-whitespace"
import "strings"
import "unicode"


func ToPropertyCase(s string) string {

	// Convert.
		result := strings.Map(
			func(r rune) rune {

				if whitespace.IsWhitespace(r) || '_' == r {
					return '-'
				} else {
					return unicode.ToLower(r)
				}
			},
			s)

	// Return
		return result
}
