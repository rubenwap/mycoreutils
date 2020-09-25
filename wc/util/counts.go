package util

import (
	"strings"
	"unicode/utf8"
)

func byteCounts(text string) int {
	return len(text)
}

func lineCounts(text string) int {
	return len(strings.Split(text, "\n"))
}

func characterCounts(text string) int {
	return utf8.RuneCountInString(text)
}

func wordCounts(text string) int {
	return len(strings.Fields(text))
}
