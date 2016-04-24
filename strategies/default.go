package strategies

import "github.com/lean-poker/poker-player-go/leanpoker"

func Default(state *leanpoker.Game) int {
	if pairRank, ok := checkPair(state); ok {
		return state.SmallBlind * pairRank
	}

	return state.SmallBlind * 2
}
