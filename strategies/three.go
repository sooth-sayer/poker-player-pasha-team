package strategies

import (
	"log"

	"github.com/lean-poker/poker-player-go/leanpoker"
)

func checkThree(game *leanpoker.Game) (int, bool) {
	player := game.Players[game.InAction]

	log.Printf("Check three %v %v", player.HoleCards, game.CommunityCards)

	count, yours := getCount(player.HoleCards, game.CommunityCards)

	if count == 4 {
		if yours {
			return player.Raise(game.SmallBlind, game.CurrentBuyIn, MaxRank), true
		}
		return player.Raise(game.SmallBlind, game.CurrentBuyIn, VeryHugeRank), true
	}

	if count == 3 {
		if yours {
			return player.Raise(game.SmallBlind, game.CurrentBuyIn, VeryHugeRank), true
		}
		return player.Raise(game.SmallBlind, game.CurrentBuyIn, HugeRank), true
	}

	return 0, false
}
