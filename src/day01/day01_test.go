package day01

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"internal/fileparser"
	"os"
	"testing"
)

func TestCalculateFuel(t *testing.T) {

	assert.Equal(t, 2, calculateFuel(12))
	assert.Equal(t, 2, calculateFuel(14))
	assert.Equal(t, 654, calculateFuel(1969))
	assert.Equal(t, 33583, calculateFuel(100756))
}

func TestCalculateFuelFileData(t *testing.T) {

	inputFile, err := os.Open("./input.txt")
	check(err)
	numbers := fileparser.ParseStringFileToIntList(inputFile)
	_ = inputFile.Close()
	total := 0
	for _, number := range numbers {
		total += calculateFuel(number)
	}
	fmt.Println(total)
}

func TestCalculateFuelInclFuelWeight(t *testing.T) {

	assert.Equal(t, 2, calculateFuelInclFuelWeight(14))
	assert.Equal(t, 966, calculateFuelInclFuelWeight(1969))
	assert.Equal(t, 50346, calculateFuelInclFuelWeight(100756))
}

func TestCalculateFuelInclFuelWeightFileData(t *testing.T) {

	inputFile, err := os.Open("./input.txt")
	check(err)
	numbers := fileparser.ParseStringFileToIntList(inputFile)
	_ = inputFile.Close()
	total := 0
	for _, number := range numbers {
		total += calculateFuelInclFuelWeight(number)
	}
	fmt.Println(total)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
