package util

import "strings"

// LeftPadToLength builds a string with the given total length padded
// with the given character
func LeftPadToLength(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

// SetStringCharAtIndex sets the character of the given string in the given index
func SetStringCharAtIndex(s string, index int, char rune) string {
	sRune := []rune(s)
	sRune[index] = char
	return string(sRune)
}
