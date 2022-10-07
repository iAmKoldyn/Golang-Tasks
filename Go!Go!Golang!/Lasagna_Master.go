package lasagna

func PreparationTime(layers []string, avgTime int) int {
    if avgTime == 0 {
        avgTime = 2
    }
    return len(layers) * avgTime
}
func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
    	switch layer {
            case "noodles":
        		noodles += 50
            case "sauce":
        		sauce += 0.2
        }
    }
	return noodles, sauce
}
func AddSecretIngredient(oldList, newList []string) []string {
    newList[len(newList)-1] = oldList[len(oldList)-1]
    return newList
}
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scalingFactor := float64(portions) / 2
    result := make([]float64, len(amounts))
    for index, amount := range amounts {
        result[index] = amount * scalingFactor
    }
	return result
}