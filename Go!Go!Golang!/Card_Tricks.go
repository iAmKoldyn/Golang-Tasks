package cards

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(slice []int, index int) int {
	if isIndexOutOfBounds(index, slice) {
		return -1
	}
	return slice[index]
}

func SetItem(slice []int, index int, value int) []int {
	if isIndexOutOfBounds(index, slice) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}
func RemoveItem(slice []int, index int) []int {
	if isIndexOutOfBounds(index, slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
func PrependItems(slice []int, newItems ...int) []int {
	return append(newItems, slice...)
}
func isIndexOutOfBounds(index int, slice []int) bool {
	return index < 0 || index >= len(slice)
}
