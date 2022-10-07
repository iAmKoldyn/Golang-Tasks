package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	var birdCount int
	for i := 0; i < len(birdsPerDay); i++ {
		birdCount += birdsPerDay[i]
	}
	return birdCount
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	var shift int
	var birdCount int
	if week == 2 {
		shift = 7
	}
	for i := 0; i < 7; i++ {
		birdCount += birdsPerDay[shift+i]
	}
	return birdCount
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
	}
	return birdsPerDay
}