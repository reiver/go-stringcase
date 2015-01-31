package stringcase


import "github.com/reiver/go-whitespace"
import "strings"
import "unicode"


// ToHeaderCase converts the string to 'Header-Case' and returns it.
func ToHeaderCase(s string) string {

	// Here we use a similar hack that the Golang strings.Title() func uses,
	// which uses the strings.Map() func but (and this is the hack'y part)
	// depends on the interation order of strings.Map().
	//
	// See: https://golang.org/src/strings/strings.go#L519
	//
	// Specifically, assumes it iterates from beginning to end.
	//
		prev := ' '
		result := strings.Map(
			func(r rune) rune {

				if whitespace.IsWhitespace(prev) || '_' == prev || '-' == prev {
					prev = r

					return unicode.ToTitle(r)
				} else if whitespace.IsWhitespace(r) || '_' == r {
					prev = r

					return '-'
				} else {
					prev = r

					return unicode.ToLower(r)

				}
			},
			s)

	// Return
		return result
}
