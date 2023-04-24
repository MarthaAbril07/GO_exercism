package main

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	x := (successRate * float64(productionRate)) / 100
	return x
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	x := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(x / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	grupos := carsCount / 10
	individuales := carsCount % 10
	return uint((grupos * 95000) + (individuales * 10000))
}
