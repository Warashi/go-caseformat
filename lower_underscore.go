package caseformat

import (
	"strings"
	"unicode"
)

type lowerUnderscore string

func (s lowerUnderscore) toCamel(toUpper bool) string {
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

func (s lowerUnderscore) ToLowerHyphen() string {
	return strings.ReplaceAll(string(s), string(underscore), string(hyphen))
}

func (s lowerUnderscore) ToLowerUnderscore() string {
	return string(s)
}

func (s lowerUnderscore) ToLowerCamel() string {
	return s.toCamel(false)
}

func (s lowerUnderscore) ToUpperHyphen() string {
	return strings.ToUpper(s.ToLowerHyphen())
}

func (s lowerUnderscore) ToUpperUnderscore() string {
	return strings.ToUpper(string(s))
}

func (s lowerUnderscore) ToUpperCamel() string {
	return s.toCamel(true)
}
