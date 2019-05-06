package caseformat

import (
	"strings"
)

type upperUnderscore string

func (s upperUnderscore) ToLowerHyphen() string {
	return strings.ToLower(s.ToUpperHyphen())
}

func (s upperUnderscore) ToLowerUnderscore() string {
	return strings.ToLower(string(s))
}

func (s upperUnderscore) ToLowerCamel() string {
	return lowerUnderscore(s.ToLowerUnderscore()).ToLowerCamel()
}

func (s upperUnderscore) ToUpperHyphen() string {
	return strings.ReplaceAll(string(s), string(underscore), string(hyphen))
}

func (s upperUnderscore) ToUpperUnderscore() string {
	return string(s)
}

func (s upperUnderscore) ToUpperCamel() string {
	return lowerUnderscore(s.ToLowerUnderscore()).ToUpperCamel()
}
