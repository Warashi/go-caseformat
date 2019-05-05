package caseformat

import (
	"strings"
	"unicode"
)

type LowerUnderscore string

func (s LowerUnderscore) toCamel(toUpper bool) string {
	var b strings.Builder
	b.Grow(len(s))

	rs := []rune(s)
	for i := range rs {
		if rs[i] == underscore {
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

func (s LowerUnderscore) ToLowerHyphen() string {
	return strings.ReplaceAll(string(s), string(underscore), string(hyphen))
}

func (s LowerUnderscore) ToLowerUnderscore() string {
	return string(s)
}

func (s LowerUnderscore) ToLowerCamel() string {
	return s.toCamel(false)
}

func (s LowerUnderscore) ToUpperHyphen() string {
	return strings.ToUpper(s.ToLowerHyphen())
}

func (s LowerUnderscore) ToUpperUnderscore() string {
	return strings.ToUpper(string(s))
}

func (s LowerUnderscore) ToUpperCamel() string {
	return s.toCamel(true)
}

