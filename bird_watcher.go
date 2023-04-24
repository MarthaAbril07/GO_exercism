package main

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	suma := 0
	for _, x := range birdsPerDay {
		suma += x
	}
	return suma
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	inicio := ((week - 1) * 7) - 1
	fin := (week * 7)
	x1 := append(birdsPerDay[inicio+1 : fin])
	return TotalBirdCount(x1)

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, x := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = x + 1
		}
	}
	return birdsPerDay
}
