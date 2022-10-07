package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0)
}

func CalculateCost(carsCount int) uint {
return uint((carsCount%10)*10000 + (carsCount/10)*95000)
    }