package util

import (
	"unicode/utf8"
)
func byteCounts(text string) int {
	return len(text)
}

// func lineCounts(text string) string {

// }

func characterCounts(text string) int {
	return utf8.RuneCountInString(text)
}

// func wordCounts(text string) string {

// }