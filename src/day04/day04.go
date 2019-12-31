package day04

import (
	"strconv"
	"strings"
)

func findNumberOfPasswords(start, end int) int {
	numberOfPasswords := 0
	for ; start <= end; start++ {
		if sixDigits(start) {
			startAsSlice := convertToSlice(start)
			valid, match := notDecreasingAndContainsDoubleDigit(startAsSlice)
			if match && valid {
				numberOfPasswords++
			}
		}
	}
	return numberOfPasswords
}

func notDecreasingAndContainsDoubleDigit(startAsSlice []string) (bool, bool) {
	valid := true
	match := false
	tmpChar := "0"

	for _, char := range startAsSlice {
		if tmpChar > char {
			valid = false
		} else if tmpChar == char {
			match = true
		}
		tmpChar = char
	}
	return valid, match
}

func sixDigits(start int) bool {
	return start > 99999 && start < 999999
}

func convertToSlice(start int) [] string {
	str := strconv.Itoa(start)
	return strings.Split(str, "")
}
