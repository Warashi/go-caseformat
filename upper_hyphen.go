package caseformat

import (
	"strings"
)

type UpperHyphen string

func (s UpperHyphen) ToLowerHyphen() string {
	return strings.ToLower(string(s))
}

func (s UpperHyphen) ToLowerUnderscore() string {
	return strings.ToLower(s.ToUpperUnderscore())
}

func (s UpperHyphen) ToLowerCamel() string {
	return LowerHyphen(s.ToLowerHyphen()).ToLowerCamel()
}

func (s UpperHyphen) ToUpperHyphen() string {
	return string(s)
}

func (s UpperHyphen) ToUpperUnderscore() string {
	return strings.ReplaceAll(string(s), string(hyphen), string(underscore))
}

func (s UpperHyphen) ToUpperCamel() string {
	return LowerHyphen(s.ToLowerHyphen()).ToUpperCamel()
}

