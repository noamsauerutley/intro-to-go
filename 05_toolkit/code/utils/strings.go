package utils

import "strings"

func MakeExcited(deflatedString string) string {
	return strings.ToUpper(deflatedString) + "!"
}
