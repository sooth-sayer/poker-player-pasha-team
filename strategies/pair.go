package strategies

import "log"
import "github.com/lean-poker/poker-player-go/leanpoker"

const (
	SmallPairRank = iota
	MiddlePairRank
	BigPairRank
)

func checkPair(state *leanpoker.Game) (int, bool) {
	player := state.Players[state.InAction]

	log.Printf("Check pair %v %v", player.HoleCards, state.CommunityCards)

	if player.HoleCards[0].Rank == player.HoleCards[1].Rank {
		log.Printf("Got own pair")
		return player.Raise(game.SmallBlind, game.CurrentBuyIn, BigPairRank), true
	}

	for _, c := range player.HoleCards {
		for _, cc := range state.CommunityCards {
			if c.Rank == cc.Rank {
				log.Printf("Got community pair")
				return player.Raise(game.SmallBlind, game.CurrentBuyIn, MiddlePairRank), true
			}
		}
	}

	return 0, false
}
