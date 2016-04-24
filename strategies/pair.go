package strategies

import "log"

const (
	SmallPairRank = iota
	MiddlePairRank
	BigPairRank
)

func checkPair(state) (int, bool) {
	player := state.Players[state.InAction]

	log.Printf("Check own pair %v %v", player.HoleCards[0].Rank, player.HoleCards[1].Rank)

	if player.HoleCards[0].Rank == player.HoleCards[1].Rank {
		return player.Raise(game.SmallBlind, game.CurrentBuyIn, BigPairRank), true
	}

	return 0, false
}
