package caseformat

import (
	"strings"
)

type upperHyphen string

func (s upperHyphen) ToLowerHyphen() string {
	return strings.ToLower(string(s))
}

func (s upperHyphen) ToLowerUnderscore() string {
	return strings.ToLower(s.ToUpperUnderscore())
}

func (s upperHyphen) ToLowerCamel() string {
	return lowerHyphen(s.ToLowerHyphen()).ToLowerCamel()
}

func (s upperHyphen) ToUpperHyphen() string {
	return string(s)
}

func (s upperHyphen) ToUpperUnderscore() string {
	return strings.ReplaceAll(string(s), string(hyphen), string(underscore))
}

func (s upperHyphen) ToUpperCamel() string {
	return lowerHyphen(s.ToLowerHyphen()).ToUpperCamel()
}
