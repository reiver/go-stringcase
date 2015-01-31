package stringcase


import "github.com/reiver/go-whitespace"
import "strings"
import "unicode"


// ToCamelCase converts the string to 'camelCase' and returns it.
func ToCamelCase(s string) string {

	// Here we use a similar hack that the Golang strings.Title() func uses,
	// which uses the strings.Map() func but (and this is the hack'y part)
	// depends on the interation order of strings.Map().
	//
	// See: https://golang.org/src/strings/strings.go#L519
	//
	// Specifically, assumes it iterates from beginning to end.
	//
		first := true
		prev := ' '
		result := strings.Map(
			func(r rune) rune {
				if first && (whitespace.IsWhitespace(prev) || '_' == prev || '-' == prev) {
					first = false
					prev = r
					return unicode.ToLower(r)
				} else if !first && (whitespace.IsWhitespace(prev) || '_' == prev || '-' == prev) {
					prev = r
					return unicode.ToTitle(r)
				} else if whitespace.IsWhitespace(r) || '_' == r || '-' == r {
					prev = r
					return -1
				} else {
					prev = r
					return unicode.ToLower(r)

				}
			},
			s)

	// Return
		return result
}
