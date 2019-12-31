package day03

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"internal/fileparser"
	"os"
	"testing"
)

func TestFindIntersectionSimpleSimple(t *testing.T) {
	assert.Equal(t, 1, findClosestIntersection([]string{"U2"}, []string{"U3"}))
}

func TestFindIntersectionSimple(t *testing.T) {
	assert.Equal(t, 3, findClosestIntersection([]string{"U3", "R3", "D4"}, []string{"R4"}))
}

func TestFindIntersection(t *testing.T) {
	assert.Equal(t, 159, findClosestIntersection([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}))
	assert.Equal(t, 135, findClosestIntersection([]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}))
}

func TestFindIntersectionFileData(t *testing.T) {

	inputFile, err := os.Open("./input.txt")
	check(err)
	firstWire, secondWire := fileparser.ParseCommaSeparatedStringToTwoStringLists(inputFile)
	_ = inputFile.Close()
	distance := findClosestIntersection(firstWire, secondWire)

	fmt.Println(distance)
}

func TestFindIntersectionWithLowestNumberOfStepsSimple(t *testing.T) {
	assert.Equal(t, 8, findIntersectionsWithLowestNumberOfSteps([]string{"U2","R2","U2"}, []string{"R2","U3"}))
}

func TestFindIntersectionWithLowestNumberOfStepsExample(t *testing.T) {
	assert.Equal(t, 610, findIntersectionsWithLowestNumberOfSteps([]string{"R75","D30","R83","U83","L12","D49","R71","U7","L72"}, []string{"U62","R66","U55","R34","D71","R55","D58","R83"}))
	assert.Equal(t, 410, findIntersectionsWithLowestNumberOfSteps([]string{"R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"}, []string{"U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"}))
}

func TestFindIntersectionWithLowestNumberOfSteps(t *testing.T) {

	inputFile, err := os.Open("./input.txt")
	check(err)
	firstWire, secondWire := fileparser.ParseCommaSeparatedStringToTwoStringLists(inputFile)
	_ = inputFile.Close()
	nrOfSteps := findIntersectionsWithLowestNumberOfSteps(firstWire, secondWire)

	fmt.Println(nrOfSteps)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
