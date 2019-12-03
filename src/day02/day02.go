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

func findNounAndVerbForOutputInIntCode(output int, state []int) (noun int, verb int) {
	nounPos := 1
	verbPos := 2
	tmpState := make([]int, len(state))
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(tmpState, state)
			tmpState[nounPos] = noun
			tmpState[verbPos] = verb
			if calculateIntCode(tmpState)[0] == output {
				return noun, verb
			}
		}
	}
	return 0, 0
}
