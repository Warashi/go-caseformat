package caseformat

import (
	"strings"
	"unicode"
)

type LowerCamel string

func (s LowerCamel) toLowerDelim(delim rune) string {
	var b strings.Builder
	b.Grow(len(s) * 2)

	rs := []rune(s)
	delimCandidate := false
	for i := range rs {
		switch {
		case delimCandidate && unicode.IsUpper(rs[i]):
			b.WriteRune(delim)
			fallthrough
		case unicode.IsUpper(rs[i]):
			b.WriteRune(unicode.ToLower(rs[i]))
			delimCandidate = false
		default:
			b.WriteRune(rs[i])
			delimCandidate = true
		}
	}
	return b.String()
}

func (s LowerCamel) ToLowerHyphen() string {
	return s.toLowerDelim(hyphen)
}

func (s LowerCamel) ToLowerUnderscore() string {
	return s.toLowerDelim(underscore)
}

func (s LowerCamel) ToLowerCamel() string {
	return string(s)
}

func (s LowerCamel) ToUpperHyphen() string {
	return strings.ToUpper(s.ToLowerHyphen())
}

func (s LowerCamel) ToUpperUnderscore() string {
	return strings.ToUpper(s.ToLowerUnderscore())
}

func (s LowerCamel) ToUpperCamel() string {
	rs := []rune(s)
	rs[0] = unicode.ToUpper(rs[0])
	return string(rs)
}
