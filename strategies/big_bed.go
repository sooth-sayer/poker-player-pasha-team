package strategies

import "log"
import "github.com/lean-poker/poker-player-go/leanpoker"

func checkBigBed(game *leanpoker.Game) bool {
	player := game.Players[game.InAction]

	result := float64(game.CurrentBuyIn) > float64(player.Stack)*0.3
	if result {
		log.Printf("Got big bet. %v %v", game.CurrentBuyIn, player.Stack)
	}
	return result
}
