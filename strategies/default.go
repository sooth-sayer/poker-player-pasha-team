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

	log.Printf("Game = %v", game)

	if rank, ok := getRank(game); ok {
		log.Printf("Make bet = %v", rank)
		return rank
	}

	if checkBigBed(game) {
		log.Printf("Got big bet")
		return 0
	}

	log.Printf("No strategy. Calling.")
	return 0
}

func getRank(game *leanpoker.Game) (int, bool) {
	cards := game.Cards()
	b := game.SmallBlind

	rank := rank_api.GetRank(cards)

	switch rank.Rank {
	case 0:
		if len(cards) < 4 {

			max := float64(0)

			switch rank.Value {
			case 11, 12, 13:
				max = 0.3
			case 14:
				max = 0.5
			default:
				max = 0
			}

			return raiseOrCall(b, game, max), true
		}
	case 1:
		log.Printf("1 %v %v", game.Cards())
		max := 0.2

		if len(cards) < 4 {
			if game.HavePair() {
				log.Printf("Have own pair %v %v", game.Cards())
				max = 0.4
			}

			switch rank.Value {
			case 11, 12, 13, 14:
				max = 0.4
			}
		}

		return raiseOrCall(b, game, max), true
	case 2:
		log.Printf("2 %v %v", game.Cards())
		return raiseOrCall(2*b, game, 0.2), true
	case 3:
		log.Printf("3 %v %v", game.Cards())
		return raiseOrCall(3*b, game, 0.3), true
	case 4:
		log.Printf("4 %v %v", game.Cards())
		return raiseOrCall(4*b, game, 0.5), true
	case 5:
		log.Printf("5 %v %v", game.Cards())
		return raiseOrCall(5*b, game, 0.5), true
	case 6:
		log.Printf("6 %v %v", game.Cards())
		return raiseOrCall(6*b, game, 1), true
	case 7:
		log.Printf("7 %v %v", game.Cards())
		return raiseOrCall(7*b, game, 1), true
	case 8:
		log.Printf("8 %v %v", game.Cards())
		return raiseOrCall(8*b, game, 1), true
	default:
		return 0, true
	}

	return 0, true
}

func raiseOrCall(bet int, game *leanpoker.Game, max float64) int {
	if game.CurrentBuyIn > bet {
		log.Printf("raiseOrCall game.CurrentBuyIn > bet, %v, %v", bet, max)
		return game.CanCall(game.CurrentBuyIn+game.SmallBlind, max)
	}

	return bet
}
