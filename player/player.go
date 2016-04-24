package player

import "github.com/lean-poker/poker-player-go/leanpoker"
import s "github.com/lean-poker/poker-player-go/strategies"

const VERSION = "Pasha Team Player 0.0.3"

func BetRequest(state *leanpoker.Game) int {
	return s.Default(state)
}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
