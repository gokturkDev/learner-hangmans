package game_loop

import (
	"fmt"
	"game_utils/player"

	"github.com/inancgumus/screen"
)

func StartGameLoop() {
	human_player := player.New()
	ai_player := player.New()
	screen.Clear()
	gameLoop(&human_player, &ai_player)
}

func gameLoop(human_player, ai_player *player.Player) {
	isPlayerTurn := true
	for { //this is the while of GO
		fmt.Printf("Your coins left: %d", human_player.GetMoney())
		fmt.Printf("AI coins left: %d", ai_player.GetMoney())
		if isPlayerTurn {
			playerTurn(human_player)
			isPlayerTurn = false
		} else {
			aiTurn(ai_player)
		}
	}
}

func playerTurn(player *player.Player) {

}

func aiTurn(player *player.Player) {

}
