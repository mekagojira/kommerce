package util

import "strings"

func IsEmptyString(str string) bool {
	return len(str) == 0
}

func IsEmptyOrWhitespace(str string) bool {
	return len(Trim(str)) == 0
}

func Trim(str string) string {
	return str[strings.Index(str, " "):]
}

func SetTrim(str *string) {
	*str = Trim(*str)
}
