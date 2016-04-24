package leanpoker

type Card struct {
	// Rank of the card. Possible values are numbers 2-10 and J,Q,K,A
	Rank string `json:"rank"`

	// Suit of the card. Possible values are: clubs,spades,hearts,diamonds
	Suit string `json:"suit"`
}

func (c *Card) IsPicture() bool {
	return c.Rank == "J" || c.Rank == "Q" || c.Rank == "K" || c.Rank == "A"
}

func (c *Card) SameRank(cc Card) bool {
	return c.Rank == cc.Rank
}

func (c *Card) IntRank() int {
	switch c.Rank {
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "10":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		return 0
	}
}
