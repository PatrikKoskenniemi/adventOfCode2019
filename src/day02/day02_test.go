package day02

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"internal/fileparser"
	"os"
	"testing"
)

func TestCalculateIntCode(t *testing.T) {

	assert.Equal(t, []int{2, 0, 0, 0, 99}, calculateIntCode([]int{1, 0, 0, 0, 99}))
	assert.Equal(t, []int{2, 3, 0, 6, 99}, calculateIntCode([]int{2, 3, 0, 3, 99}))
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801}, calculateIntCode([]int{2, 4, 4, 5, 99, 0}))
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, calculateIntCode([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}))
}

func TestCalculateIntCodeFileData(t *testing.T) {

	inputFile, err := os.Open("./input.txt")
	check(err)
	state := fileparser.ParseCommaSeparatedStringToIntList(inputFile)
	_ = inputFile.Close()
	state[1] = 12
	state[2] = 2
	finalState := calculateIntCode(state)

	fmt.Println(finalState[0])
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
