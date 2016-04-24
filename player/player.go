package player

import "github.com/lean-poker/poker-player-go/leanpoker"

const VERSION = "Pasha Team Player 0.0.1"

func BetRequest(state *leanpoker.Game) int {
	return 1
}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
