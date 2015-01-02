package stringutil

import (
	"bytes"
	"unicode"
)

// ToCamelCase returns a copy of the string s with all Unicode letters mapped to their camel case.
// It will convert to upper case previous letter of '_' and first letter, and remove letter of '_'.
func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}
	result := make([]rune, 0, len(s))
	upper := false
	for _, r := range s {
		if r == '_' {
			upper = true
			continue
		}
		if upper {
			result = append(result, unicode.ToUpper(r))
			upper = false
			continue
		}
		result = append(result, r)
	}
	result[0] = unicode.ToUpper(result[0])
	return string(result)
}

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
