package caseformat

import (
	"strings"
	"unicode"
)

type lowerHyphen string

func (s lowerHyphen) toCamel(toUpper bool) string {
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

func (s lowerHyphen) ToLowerHyphen() string {
	return string(s)
}

func (s lowerHyphen) ToLowerUnderscore() string {
	return strings.ReplaceAll(string(s), string(hyphen), string(underscore))
}

func (s lowerHyphen) ToLowerCamel() string {
	return s.toCamel(false)
}

func (s lowerHyphen) ToUpperHyphen() string {
	return strings.ToUpper(string(s))
}

func (s lowerHyphen) ToUpperUnderscore() string {
	return strings.ToUpper(s.ToLowerUnderscore())
}

func (s lowerHyphen) ToUpperCamel() string {
	return s.toCamel(true)
}

