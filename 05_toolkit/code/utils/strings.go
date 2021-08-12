package utils

import "strings"

//transforms a sentence to all caps with exclamation
func MakeExcited(sentence string) string  {
	return strings.ToUpper(sentence) + "!"
}