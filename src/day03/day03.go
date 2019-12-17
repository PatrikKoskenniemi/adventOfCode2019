package day03

import (
	"math"
	"strconv"
)

type position struct {
	x int
	y int
}

func findIntersection(firstWire, secondWire []string) int {
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

	return findClosestIntersection(firstWirePositions, secondWirePositions)
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
func findClosestIntersection(firstWire []position, secondWire []position) int {
	shortestDistance := math.MaxInt32
	for _, firstWireElem := range firstWire {
		for _, secondWireElem := range secondWire {
			if firstWireElem == secondWireElem {
				distance := manhattanDistance(0, 0, firstWireElem.x, firstWireElem.y)
				if distance < shortestDistance {
					shortestDistance = distance
				}
			}
		}
	}
	return shortestDistance
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
