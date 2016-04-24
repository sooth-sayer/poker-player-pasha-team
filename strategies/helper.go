package strategies

import (
	"github.com/lean-poker/poker-player-go/leanpoker"
)

func getCount(your, board []leanpoker.Card) (int, bool) {
	yours := your[0].SameRank(your[1])

	cards := make([]leanpoker.Card, 0)
	cards = append(cards, your...)
	cards = append(cards, board...)

	max := 0

	for _, c := range cards {
		count := -1

		for _, cc := range cards {
			if c.SameRank(cc) {
				count = count + 1
			}
		}

		if max < count {
			max = count
		}
	}

	return max, yours
}
