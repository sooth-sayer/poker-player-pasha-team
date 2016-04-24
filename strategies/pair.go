package strategies

import "log"
import "github.com/lean-poker/poker-player-go/leanpoker"

const (
	SmallPairRank = iota
	MiddlePairRank
	BigPairRank
	VeryBigPairRank
	HugePairRank
)

func checkPair(game *leanpoker.Game) (int, bool) {
	player := game.Players[game.InAction]

	log.Printf("Check pair %v %v", player.HoleCards, game.CommunityCards)

	if player.HoleCards[0].Rank == player.HoleCards[1].Rank {
		log.Printf("Got own pair")

		if player.HoleCards[0].IsPicture() {
			return player.Raise(game.SmallBlind, game.CurrentBuyIn, HugePairRank), true
		}

		return player.Raise(game.SmallBlind, game.CurrentBuyIn, BigPairRank), true
	}

	for _, c := range player.HoleCards {
		for _, cc := range game.CommunityCards {
			if c.Rank == cc.Rank {
				log.Printf("Got community pair")

				if player.HoleCards[0].IsPicture() {
					return player.Raise(game.SmallBlind, game.CurrentBuyIn, VeryBigPairRank), true
				}

				return player.Raise(game.SmallBlind, game.CurrentBuyIn, MiddlePairRank), true
			}
		}
	}

	return 0, false
}
