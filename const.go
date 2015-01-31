package stringcase


import "github.com/reiver/go-whitespace"
import "strings"
import "unicode"


// ToConstCase converts the string to 'CONST_CASE' and returns it.
func ToConstCase(s string) string {

	// Convert.
		result := strings.Map(
			func(r rune) rune {

				if whitespace.IsWhitespace(r) || '-' == r {
					return '_'
				} else {
					return unicode.ToUpper(r)
				}
			},
			s)

	// Return
		return result
}
