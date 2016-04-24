package strategies

import (
	"log"

	"github.com/lean-poker/poker-player-go/leanpoker"
)

func Default(game *leanpoker.Game) int {
	if pairRank, ok := checkPair(game); ok {
		return game.SmallBlind * pairRank
	}

	if checkBigBed(game) {
		log.Printf("Got big bet")
		return 0
	}

	return game.SmallBlind * 2
}
