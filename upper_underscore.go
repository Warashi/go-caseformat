package caseformat

import (
	"strings"
)

type UpperUnderscore string

func (s UpperUnderscore) ToLowerHyphen() string {
	return strings.ToLower(s.ToUpperHyphen())
}

func (s UpperUnderscore) ToLowerUnderscore() string {
	return strings.ToLower(string(s))
}

func (s UpperUnderscore) ToLowerCamel() string {
	return LowerUnderscore(s.ToLowerUnderscore()).ToLowerCamel()
}

func (s UpperUnderscore) ToUpperHyphen() string {
	return strings.ReplaceAll(string(s), string(underscore), string(hyphen))
}

func (s UpperUnderscore) ToUpperUnderscore() string {
	return string(s)
}

func (s UpperUnderscore) ToUpperCamel() string {
	return LowerUnderscore(s.ToLowerUnderscore()).ToUpperCamel()
}

