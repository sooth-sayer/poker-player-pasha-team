package strategies

import (
	"log"

	"github.com/lean-poker/poker-player-go/leanpoker"
	"github.com/lean-poker/poker-player-go/rank_api"
)

const (
	BaseRank = 10
)

func Default(game *leanpoker.Game) int {
	if rank, ok := getRank(game); ok {
		log.Printf("Make bet = %v", rank)
		return rank
	}

	if checkBigBed(game) {
		log.Printf("Got big bet")
		return 0
	}

	log.Printf("No strategy. Calling.")
	return game.Call()
}

func getRank(game *leanpoker.Game) (int, bool) {
	cards := game.Cards()

	if len(cards) < 5 {
		log.Printf("Less 5 cards")
		return 0, false
	}

	rank := rank_api.GetRank(cards)

	b := game.SmallBlind

	switch rank.Rank {
	case 1:
		return BaseRank, true
	case 2:
		return raiseOrCall(2*b, game, 0.1), true
	case 3:
		return raiseOrCall(3*b, game, 0.3), true
	case 4:
		return raiseOrCall(4*b, game, 0.5), true
	case 5:
		return raiseOrCall(5*b, game, 0.5), true
	case 6:
		return raiseOrCall(6*b, game, 1), true
	case 7:
		return raiseOrCall(7*b, game, 1), true
	case 8:
		return raiseOrCall(8*b, game, 1), true
	default:
		return 0, true
	}
}

func raiseOrCall(bet int, game *leanpoker.Game, max float64) int {
	if game.CurrentBuyIn > bet {
		return game.CanCall(game.CurrentBuyIn+game.SmallBlind, max)
	}

	return bet
}
