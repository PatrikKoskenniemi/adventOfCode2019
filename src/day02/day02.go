package day02

func calculateIntCode(state []int) []int {

	instructionPointer := 0
	for state[instructionPointer] != 99 {
		firstParam := state[instructionPointer+1]
		secondParam := state[instructionPointer+2]
		resultPos := state[instructionPointer+3]
		switch state[instructionPointer] {
		case 1:
			state[resultPos] = add(state[firstParam], state[secondParam])
		case 2:
			state[resultPos] = multiply(state[firstParam], state[secondParam])
		}
		instructionPointer += 4
	}
	return state
}

func add(first, second int) int {
	return first + second
}

func multiply(first, second int) int {
	return first * second
}
