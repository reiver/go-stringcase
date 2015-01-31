package stringcase


import "github.com/reiver/go-whitespace"
import "strings"
import "unicode"


// ToSnakeCase converts the string to 'snake_case' and returns it.
func ToSnakeCase(s string) string {

	// Convert.
		result := strings.Map(
			func(r rune) rune {

				if whitespace.IsWhitespace(r) || '-' == r {
					return '_'
				} else {
					return unicode.ToLower(r)
				}
			},
			s)

	// Return
		return result
}
