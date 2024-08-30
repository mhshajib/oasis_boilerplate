package utils

import (
	"strings"
	"unicode"
)

// Converts any string format to camelCase
func toCamelCase(str string) string {
	runes := []rune(str)
	var result []rune
	isFirst := true
	nextUpper := false

	for _, r := range runes {
		if r == '_' || r == '-' || unicode.IsSpace(r) {
			nextUpper = true
			continue
		}

		if isFirst {
			result = append(result, unicode.ToLower(r))
			isFirst = false
		} else if nextUpper {
			result = append(result, unicode.ToUpper(r))
			nextUpper = false
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// Converts any string format to TitleCase (every word starts with an uppercase letter)
func toTitleCase(str string) string {
	runes := []rune(str)
	var result []rune
	nextUpper := true

	for i, r := range runes {
		if r == '_' || r == '-' || unicode.IsSpace(r) {
			nextUpper = true
			continue
		}

		if nextUpper {
			result = append(result, unicode.ToUpper(r))
			nextUpper = false
		} else {
			result = append(result, runes[i])
		}
	}
	return string(result)
}

// Converts any string format to snake_case
func toSnakeCase(str string) string {
	runes := []rune(str)
	var result []rune
	previousWasUnderscore := false

	for i, r := range runes {
		lowerR := unicode.ToLower(r)
		if r == '-' || unicode.IsSpace(r) {
			if !previousWasUnderscore && len(result) > 0 {
				result = append(result, '_')
				previousWasUnderscore = true
			}
		} else if unicode.IsUpper(r) && i > 0 && !previousWasUnderscore {
			result = append(result, '_')
			result = append(result, lowerR)
			previousWasUnderscore = false
		} else {
			result = append(result, lowerR)
			previousWasUnderscore = lowerR == '_'
		}
	}
	return strings.Trim(string(result), "_")
}

func whitespaceToSnakeCase(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return '_' // Replace whitespace with underscore
		}
		return r
	}, str)
}

// ProcessString processes the string based on given rules
func ProcessString(input string) (string, string, string) {
	processedInput := whitespaceToSnakeCase(input)
	camelCase := toCamelCase(processedInput)
	titleCase := toTitleCase(processedInput)
	snakeCase := toSnakeCase(processedInput)

	return titleCase, snakeCase, camelCase
}
