package strategies

import "log"

func checkBigBed(game *leanpoker.Godeps) bool {
	player := game.Players[game.InAction]

	result := game.CurrentBuyIn > player.Stack*0.3
	if result {
		log.Printf("Got big bet. %v %v", game.CurrentBuyIn, player.Stack)
	}
	return result
}
