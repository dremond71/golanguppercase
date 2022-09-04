package golanguppercase

import (
	"strings"
)

//
// Uppercase function uppercases a string
//
func Uppercase(someValue string) string {
	return strings.ToUpper(someValue)
}
