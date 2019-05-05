package caseformat

import (
	"unicode"
)

type UpperCamel string

func (s UpperCamel) ToLowerHyphen() string {
	return LowerCamel(s.ToLowerCamel()).ToLowerHyphen()
}

func (s UpperCamel) ToLowerUnderscore() string {
	return LowerCamel(s.ToLowerCamel()).ToLowerUnderscore()
}

func (s UpperCamel) ToLowerCamel() string {
	rs := []rune(s)
	rs[0] = unicode.ToLower(rs[0])
	return string(rs)
}

func (s UpperCamel) ToUpperHyphen() string {
	return LowerCamel(s.ToLowerCamel()).ToUpperHyphen()
}

func (s UpperCamel) ToUpperUnderscore() string {
	return LowerCamel(s.ToLowerCamel()).ToUpperUnderscore()
}

func (s UpperCamel) ToUpperCamel() string {
	return string(s)
}


