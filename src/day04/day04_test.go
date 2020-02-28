package day04

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumberOfPasswordsSixDigits(t *testing.T) {
	assert.Equal(t, 0, findNumberOfPasswords(0,99999))
}

func TestFindNumberOfPasswordsAdjacentDigitsValid (t *testing.T) {
	assert.Equal(t, 2, findNumberOfPasswords(122345,122346))
}

func TestFindNumberOfPasswordsAdjacentDigitsInValid (t *testing.T) {
	assert.Equal(t, 0, findNumberOfPasswords(122350,122354))
}

func TestFindNumberOfPasswordsNoAdjacentDigits (t *testing.T) {
	assert.Equal(t, 0, findNumberOfPasswords(123456,123457))
}

func TestFindNumberOfPasswords(t *testing.T) {
	numberOfPasswords := findNumberOfPasswords(245182, 790572)
	fmt.Println(numberOfPasswords)
}

func TestFindNumberOfPasswordsNotPartOfLargerGroup (t *testing.T) {
	assert.Equal(t, 1, findNumberOfPasswords(123444,123445))
}

func TestFindNumberOfPasswordsNotPartOfLargerGroup2 (t *testing.T) {
	assert.Equal(t, 1, findNumberOfPasswords(111122,111122))
}

func TestFindNumberOfPasswordsNotPartOfLargerGroup3 (t *testing.T) {
	assert.Equal(t, 1, findNumberOfPasswords(112223,112223))
}
func TestFindNumberOfPasswordsNotPartOfLargerGroup4 (t *testing.T) {
	assert.Equal(t, 0, findNumberOfPasswords(789999,789999))
}
