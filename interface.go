package caseformat

type CaseFormat interface {
	ToLowerHyphen() string
	ToLowerUnderscore() string
	ToLowerCamel() string
	ToUpperHyphen() string
	ToUpperUnderscore() string
	ToUpperCamel() string
}

func FromLowerHyphen(s string) CaseFormat {
	return lowerHyphen(s)
}

func FromLowerUnderscore(s string) CaseFormat {
	return lowerUnderscore(s)
}

func FromLowerCamel(s string) CaseFormat {
	return lowerCamel(s)
}

func FromUpperHyphen(s string) CaseFormat {
	return upperHyphen(s)
}

func FromUpperUnderscore(s string) CaseFormat {
	return upperUnderscore(s)
}

func FromUpperCamel(s string) CaseFormat {
	return upperCamel(s)
}
