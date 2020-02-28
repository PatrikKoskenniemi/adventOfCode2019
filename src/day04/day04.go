package day04

import (
	"fmt"
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
				fmt.Println(start)
				numberOfPasswords++
			}
		}
	}
	return numberOfPasswords
}

func notDecreasingAndContainsDoubleDigit(startAsSlice []string) (bool, bool) {
	valid := true
	prevChar := "0"
	for _, currChar := range startAsSlice {
		if prevChar > currChar {
			valid = false
		}
		prevChar = currChar
	}

	match := false
	prevChar = "0"
	matchingChar := ""

	/*for _, currChar := range startAsSlice {
		if prevChar == currChar {
			if match && matchingChar != currChar {
				return valid, match
			} else if match && matchingChar == currChar {
				match = false
			} else{
				matchingChar = currChar
				match = true
			}
		}
		prevChar = currChar
	}*/

	for _, currChar := range startAsSlice {
		if matchingChar == currChar && prevChar == currChar {
			match = false
		} else if match && prevChar != currChar {
			return valid, match
		} else if prevChar == currChar {
			match = true
			matchingChar = currChar
		}
		prevChar = currChar
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
