package caseformat

import (
	"unicode"
)

type upperCamel string

func (s upperCamel) ToLowerHyphen() string {
	return lowerCamel(s.ToLowerCamel()).ToLowerHyphen()
}

func (s upperCamel) ToLowerUnderscore() string {
	return lowerCamel(s.ToLowerCamel()).ToLowerUnderscore()
}

func (s upperCamel) ToLowerCamel() string {
	rs := []rune(s)
	rs[0] = unicode.ToLower(rs[0])
	return string(rs)
}

func (s upperCamel) ToUpperHyphen() string {
	return lowerCamel(s.ToLowerCamel()).ToUpperHyphen()
}

func (s upperCamel) ToUpperUnderscore() string {
	return lowerCamel(s.ToLowerCamel()).ToUpperUnderscore()
}

func (s upperCamel) ToUpperCamel() string {
	return string(s)
}


