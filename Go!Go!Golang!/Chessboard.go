package chessboard

type File []bool
type Chessboard map[string]File

func CountInRank(cb Chessboard, rank int) (ret int) {
	if rank > 8 || rank < 1 {
		return 0
	}
	for _, file := range cb {
		if file[rank-1] {
			ret += 1
		}
	}
	return ret
}

func CountInFile(cb Chessboard, file string) (ret int) {
	for _, rankOccupiedOn := range cb[file] {
		if rankOccupiedOn {
			ret += 1
		}
	}
	return ret
}

func CountAll(cb Chessboard) (ret int) {
	for _, fileSquares := range cb {
		ret += len(fileSquares)
	}
	return ret
}

func CountOccupied(cb Chessboard) (ret int) {
	for i := 65; i < 73; i++ {
		ret += CountInFile(cb, string(rune(i)))
	}
	return ret
}
