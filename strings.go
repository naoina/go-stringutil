package stringutil

import (
	"bytes"
	"unicode"
)

// ToUpperCamelCase returns a copy of the string s with all Unicode letters mapped to their camel case.
// It will convert to upper case previous letter of '_' and first letter, and remove letter of '_'.
func ToUpperCamelCase(s string) string {
	if s == "" {
		return ""
	}
	upper := true
	var result bytes.Buffer
	for _, c := range s {
		if c == '_' {
			upper = true
			continue
		}
		if upper {
			result.WriteRune(unicode.ToUpper(c))
			upper = false
			continue
		}
		result.WriteRune(c)
	}
	return result.String()
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

// ToSnakeCaseASCII is similar to ToSnakeCase, but optimized for only the ASCII
// characters.
// ToSnakeCaseASCII is faster than ToSnakeCase, but doesn't work correctly if
// contains non-ASCII characters.
func ToSnakeCaseASCII(s string) string {
	if s == "" {
		return ""
	}
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isUpperASCII(c) {
			result = append(result, '_')
		}
		result = append(result, toLowerASCII(c))
	}
	if result[0] == '_' {
		return string(result[1:])
	}
	return string(result)
}

func isUpperASCII(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func toLowerASCII(c byte) byte {
	if isUpperASCII(c) {
		return c + 'a' - 'A'
	}
	return c
}
