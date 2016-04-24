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
