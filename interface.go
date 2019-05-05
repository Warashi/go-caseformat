package caseformat

type CaseFormat interface {
	ToLowerHyphen() string
	ToLowerUnderscore() string
	ToLowerCamel() string
	ToUpperHyphen() string
	ToUpperUnderscore() string
	ToUpperCamel() string
}
