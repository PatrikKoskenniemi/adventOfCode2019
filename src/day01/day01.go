package day01

func calculateFuel(weight int) int {
	return weight/3 - 2
}

func calculateFuelInclFuelWeight(weight int) int {
	fuel := calculateFuel(weight)

	if fuel <= 0 {
		return 0
	}
	return fuel + calculateFuelInclFuelWeight(fuel)
}
