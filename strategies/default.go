package strategies

import "github.com/lean-poker/poker-player-go/leanpoker"

func Default(game *leanpoker.Game) int {
	if pairRank, ok := checkPair(game); ok {
		return game.SmallBlind * pairRank
	}

	return game.SmallBlind * 2
}
