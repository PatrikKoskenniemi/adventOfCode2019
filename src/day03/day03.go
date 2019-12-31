package day03

import (
	"math"
	"strconv"
)

type position struct {
	x int
	y int
}

func findClosestIntersection(firstWire, secondWire []string) int {
	firstWirePositions, secondWirePositions := findPositions(firstWire, secondWire)
	intersections := findIntersections(firstWirePositions, secondWirePositions)
	return calculateClosestIntersection(intersections)
}

func findPositions(firstWire, secondWire []string) ([]position, []position) {
	firstWirePositions := make([]position, 0)
	secondWirePositions := make([]position, 0)
	currX, currY := 0, 0
	for _, instruction := range firstWire {
		currX, currY = calculateSteps(instruction, currX, currY, &firstWirePositions)
	}
	currX, currY = 0, 0
	for _, instruction := range secondWire {
		currX, currY = calculateSteps(instruction, currX, currY, &secondWirePositions)
	}
	return firstWirePositions, secondWirePositions
}

func calculateSteps(instruction string, x, y int, wire *[]position) (int, int) {
	switch instruction[0] {
	case 'U':
		x, y = moveUp(wire, x, y, instruction[1:])
	case 'R':
		x, y = moveRight(wire, x, y, instruction[1:])
	case 'D':
		x, y = moveDown(wire, x, y, instruction[1:])
	case 'L':
		x, y = moveLeft(wire, x, y, instruction[1:])
	}
	return x, y
}

func moveUp(wire *[]position, x, y int, stepString string) (int, int) {
	steps, _ := strconv.Atoi(stepString)
	for i := 0; i < steps; i++ {
		y++
		*wire = append(*wire, position{x, y})
	}
	return x, y
}
func moveRight(wire *[]position, x, y int, stepString string) (int, int) {
	steps, _ := strconv.Atoi(stepString)
	for ; steps > 0; steps-- {
		x++
		*wire = append(*wire, position{x, y})
	}
	return x, y
}
func moveDown(wire *[]position, x, y int, stepString string) (int, int) {
	steps, _ := strconv.Atoi(stepString)
	for ; steps > 0; steps-- {
		y--
		*wire = append(*wire, position{x, y})
	}
	return x, y
}
func moveLeft(wire *[]position, x, y int, stepString string) (int, int) {
	steps, _ := strconv.Atoi(stepString)
	for ; steps > 0; steps-- {
		x--
		*wire = append(*wire, position{x, y})
	}
	return x, y
}
func calculateClosestIntersection(intersections []position) int {
	shortestDistance := math.MaxInt32
	for _, intersection := range intersections {
		distance := manhattanDistance(0, 0, intersection.x, intersection.y)
		if distance < shortestDistance {
			shortestDistance = distance
		}
	}
	return shortestDistance
}

func findIntersections(firstWire []position, secondWire []position) []position {
	intersections := make([]position, 0)
	for _, firstWireElem := range firstWire {
		for _, secondWireElem := range secondWire {
			if firstWireElem == secondWireElem {
				intersections = append(intersections, firstWireElem)
			}
		}
	}
	return intersections
}

func manhattanDistance(fromX, fromY, toX, toY int) int {
	return Abs(fromX-toX) + Abs(fromY-toY)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findIntersectionsWithLowestNumberOfSteps(firstWire, secondWire []string) int {
	firstWirePositions, secondWirePositions := findPositions(firstWire, secondWire)
	intersections := findIntersections(firstWirePositions, secondWirePositions)
	stepsForWiresToIntersections := calculateNumberOfStepsForWiresToIntersections(intersections, firstWirePositions, secondWirePositions)
	shortestDistance := math.MaxInt32
	for _, totSteps := range stepsForWiresToIntersections {
		if shortestDistance > totSteps {
			shortestDistance = totSteps
		}
	}
	return shortestDistance
}

func calculateNumberOfStepsForWiresToIntersections(intersections, firstWirePositions, secondWirePositions []position) []int {
	nrOfStepsPerIntersection := make([]int, len(intersections))
	for index, intersection := range intersections {
		nrOfStepsPerIntersection[index] = calculateStepsForWiresToIntersection(intersection, firstWirePositions, secondWirePositions)
	}
	return nrOfStepsPerIntersection
}

func calculateStepsForWiresToIntersection(intersection position, firstWirePositions, secondWirePositions []position) int {
	i := 0
	totalNumberOfSteps := 1
	for firstWirePositions[i] != intersection {
		totalNumberOfSteps++
		i++
	}
	i = 0
	totalNumberOfSteps++
	for secondWirePositions[i] != intersection {
		totalNumberOfSteps++
		i++
	}
	return totalNumberOfSteps
}
