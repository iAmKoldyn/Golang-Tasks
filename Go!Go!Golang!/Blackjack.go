package blackjack

func ParseCard(card string) int {
	switch card {
        case "ace": return 11
    	case "two": return 2
    	case "three": return 3
    	case "four": return 4
    	case "five": return 5
    	case "six": return 6
    	case "seven": return 7
    	case "eight": return 8
    	case "nine": return 9
    	case "ten": return 10
    	case "jack": return 10
    	case "queen": return 10
    	case "king": return 10
    	default: return 0
    }
}

func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1) + ParseCard(card2) == 21
}

func LargeHand(isBlackjack bool, dealerScore int) string {
	if (!isBlackjack) {

		return "P"
    }
	if (isBlackjack && dealerScore < 10) {

		return "W"
    }

	return "S"
}

func SmallHand(handScore int, dealerScore int) string {
	if (handScore >= 17) {
		return "S"
    } else if (handScore <= 11) {
		return "H"
    } else if (dealerScore >= 7) {
		return "H"
    } else {
		return "S"
    }
}

func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)
	if 20 < handScore {
		return LargeHand(IsBlackjack(card1, card2), dealerScore)
	}
	return SmallHand(handScore, dealerScore)
}
