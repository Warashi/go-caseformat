package caseformat

import (
	"strings"
	"unicode"
)

type lowerCamel string

type beforeCase int

const (
	none beforeCase = iota
	lower
	upper
	deliminator
)

func (s lowerCamel) toLowerDelim(delim rune) string {
	rs := []rune(s)
	ret := make([]rune, 0, len(rs)*2)

	before := none
	for i := len(rs) - 1; i >= 0; i-- {
		if before == none {
			ret = append(ret, unicode.ToLower(rs[i]))
		}
		switch {
		case unicode.IsLower(rs[i]):
			switch before {
			case lower, deliminator:
				ret = append(ret, rs[i])
			case upper:
				ret = append(ret, delim, rs[i])
			}
			before = lower
		case unicode.IsUpper(rs[i]):
			switch before {
			case lower:
				ret = append(ret, unicode.ToLower(rs[i]), delim)
				before = deliminator
				continue
			case upper, deliminator:
				ret = append(ret, unicode.ToLower(rs[i]))
			}
			before = upper
		}
	}

	// reverse ret
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret)
}

func (s lowerCamel) ToLowerHyphen() string {
	return s.toLowerDelim(hyphen)
}

func (s lowerCamel) ToLowerUnderscore() string {
	return s.toLowerDelim(underscore)
}

func (s lowerCamel) ToLowerCamel() string {
	return string(s)
}

func (s lowerCamel) ToUpperHyphen() string {
	return strings.ToUpper(s.ToLowerHyphen())
}

func (s lowerCamel) ToUpperUnderscore() string {
	return strings.ToUpper(s.ToLowerUnderscore())
}

func (s lowerCamel) ToUpperCamel() string {
	rs := []rune(s)
	rs[0] = unicode.ToUpper(rs[0])
	return string(rs)
}
