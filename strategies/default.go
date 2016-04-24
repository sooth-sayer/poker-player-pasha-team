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

	if len(cards) < 3 {
		max := float64(0)
		if game.HavePair() {
			max = 0.3
			return raiseOrCall(b, game, max), true
		}
		if game.IsPictures() {
			max = 0.1
			return raiseOrCall(b, game, max), true
		}

		return 0, true
	}

	switch rank.Rank {
	case 0:
		log.Printf("0 %v %v", game.Cards())

		if len(cards) < 6 {

			max := float64(0)

			switch rank.Value {
			case 11, 12, 13:
				max = 0.1
			case 14:
				max = 0.1
			default:
				max = 0
			}

			return raiseOrCall(b, game, max), true
		}

		return 0, true
	case 1:
		log.Printf("1 %v %v", game.Cards())
		max := 0.5

		if game.HavePair() {
			log.Printf("Have own pair %v %v", game.Cards())
			max = 0.7
		}

		switch rank.Value {
		case 11, 12, 13, 14:
			max = 0.6
		}

		player := game.Players[game.InAction]

		if int(rank.Value) != player.HoleCards[0].IntRank() && int(rank.Value) != player.HoleCards[1].IntRank() {
			return raiseOrCall(b, game, 0.01), true
		}

		return raiseOrCall(10*b, game, max), true
	case 2:
		log.Printf("2 %v %v", game.Cards())
		return raiseOrCall(10*b, game, 1), true
	case 3:
		log.Printf("3 %v %v", game.Cards())
		return raiseOrCall(20*b, game, 1), true
	case 4:
		log.Printf("4 %v %v", game.Cards())
		return raiseOrCall(20*b, game, 1), true
	case 5:
		log.Printf("5 %v %v", game.Cards())
		return raiseOrCall(15*b, game, 1), true
	case 6:
		log.Printf("6 %v %v", game.Cards())
		return raiseOrCall(20*b, game, 1), true
	case 7:
		log.Printf("7 %v %v", game.Cards())
		return raiseOrCall(100*b, game, 1), true
	case 8:
		log.Printf("8 %v %v", game.Cards())
		return raiseOrCall(100*b, game, 1), true
	default:
		return 0, true
	}

	return 0, true
}

func raiseOrCall(bet int, game *leanpoker.Game, max float64) int {
	if game.CurrentBuyIn > bet {
		log.Printf("raiseOrCall game.CurrentBuyIn > bet, %v, %v", bet, max)
		return game.CanCall(game.CurrentBuyIn, max)
	}

	return bet
}
