package player

import "github.com/lean-poker/poker-player-go/leanpoker"
import s "github.com/lean-poker/poker-player-go/strategies"

const VERSION = "Pasha Team Player 0.0.3"

func BetRequest(game *leanpoker.Game) int {
	return s.Default(game)
}

func Showdown(game *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
