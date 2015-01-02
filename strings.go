package stringutil

import (
	"bytes"
	"unicode"
)

// ToSnakeCase returns a copy of the string s with all Unicode letters mapped to their snake case.
// It will insert letter of '_' at position of previous letter of uppercase and all
// letters convert to lower case.
func ToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	var result bytes.Buffer
	for _, c := range s {
		if unicode.IsUpper(c) {
			result.WriteByte('_')
		}
		result.WriteRune(unicode.ToLower(c))
	}
	s = result.String()
	if s[0] == '_' {
		return s[1:]
	}
	return s
}
