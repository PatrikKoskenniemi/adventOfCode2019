package fileparser

import (
	"bufio"
	"os"
	"strconv"
)

func ParseStringFileToIntList(file *os.File) []int {

	var numbers []int
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		number, _ := strconv.Atoi(fileScanner.Text())
		numbers = append(numbers, number)
	}

	return numbers

}
