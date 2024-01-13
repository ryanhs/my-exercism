package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	max := len(birdsPerDay)

	for i := 0; i < max; i++ {
		total += birdsPerDay[i]
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := (week - 1) * 7
	endIndex := startIndex + 7

	// week not found
	if len(birdsPerDay) < startIndex || len(birdsPerDay) < endIndex {
		return 0
	}

	return TotalBirdCount(birdsPerDay[startIndex:endIndex])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	maxDay := len(birdsPerDay)
	for i := 0; i < maxDay; i++ {
		newValue := birdsPerDay[i]

		// on every second day
		if i == 0 || i%2 == 0 {
			newValue += 1
		}

		birdsPerDay[i] = newValue
	}

	return birdsPerDay
}
