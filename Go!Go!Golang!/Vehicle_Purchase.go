package purchase

func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

func ChooseVehicle(option1, option2 string) string {
	return min(option1, option2) + " is clearly the better choice."
}
func min(first string, second string) string {
	if first < second {
		return first
	}
	return second
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return .8 * originalPrice
	} else if age < 10 {
		return .7 * originalPrice
	} else {
		return .5 * originalPrice
	}
}
