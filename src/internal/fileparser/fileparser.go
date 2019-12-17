package fileparser

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

func ParseOneStringPerRowToIntList(file *os.File) []int {

	var numbers []int
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		number, _ := strconv.Atoi(fileScanner.Text())
		numbers = append(numbers, number)
	}

	return numbers
}

func ParseCommaSeparatedStringToIntList(file *os.File) []int {

	// An anonymous function declaration to avoid repeating main()
	ScanCSV := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		commaidx := bytes.IndexByte(data, ',')
		if commaidx > 0 {
			// we need to return the next position
			buffer := data[:commaidx]
			return commaidx + 1, bytes.TrimSpace(buffer), nil
		}

		// if we are at the end of the string, just return the entire buffer
		if atEOF {
			// but only do that when there is some data. If not, this might mean
			// that we've reached the end of our input CSV string
			if len(data) > 0 {
				return len(data), bytes.TrimSpace(data), nil
			}
		}

		// when 0, nil, nil is returned, this is a signal to the interface to read
		// more data in from the input reader. In this case, this input is our
		// string reader and this pretty much will never occur.
		return 0, nil, nil
	}

	var numbers []int
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(ScanCSV)
	for fileScanner.Scan() {
		number, _ := strconv.Atoi(fileScanner.Text())
		numbers = append(numbers, number)
	}

	return numbers

}

func ParseCommaSeparatedStringToTwoStringLists(file *os.File) ([]string, []string) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()
	firstRow := fileScanner.Text()
	fileScanner.Scan()
	secondRow := fileScanner.Text()

	return strings.Split(firstRow, ","), strings.Split(secondRow, ",")
}
