package caseformat

import (
	"strings"
	"unicode"
)

type LowerHyphen string

func (s LowerHyphen) toCamel(toUpper bool) string {
	var b strings.Builder
	b.Grow(len(s))

	rs := []rune(s)
	for i := range rs {
		if rs[i] == hyphen {
			toUpper = true
			continue
		}
		if toUpper {
			b.WriteRune(unicode.ToUpper(rs[i]))
			toUpper = false
			continue
		}
		b.WriteRune(rs[i])
	}
	return b.String()
}

func (s LowerHyphen) ToLowerHyphen() string {
	return string(s)
}

func (s LowerHyphen) ToLowerUnderscore() string {
	return strings.ReplaceAll(string(s), string(hyphen), string(underscore))
}

func (s LowerHyphen) ToLowerCamel() string {
	return s.toCamel(false)
}

func (s LowerHyphen) ToUpperHyphen() string {
	return strings.ToUpper(string(s))
}

func (s LowerHyphen) ToUpperUnderscore() string {
	return strings.ToUpper(s.ToLowerUnderscore())
}

func (s LowerHyphen) ToUpperCamel() string {
	return s.toCamel(true)
}

