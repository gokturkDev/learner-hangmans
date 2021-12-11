package main

import (
	"fmt"
	"game_utils/player"
	"os"
	"os/exec"

	"github.com/inancgumus/screen"
)

func main() {
	fmt.Print("Welcome to the GOLang coin flipper-uh!\nPress enter to continue:")
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	StartGameLoop()
}

func StartGameLoop() {
	human_player := player.New()
	ai_player := player.New()
	gameLoop(&human_player, &ai_player)
}

func gameLoop(human_player, ai_player *player.Player) {
	isPlayerTurn := true
	for { //this is the while of GO
		screen.Clear()
		fmt.Printf("Your coins left: %d \n", human_player.GetMoney())
		fmt.Printf("AI coins left: %d \n", ai_player.GetMoney())
		if isPlayerTurn {
			playerTurn(human_player)
			isPlayerTurn = false
		} else {
			aiTurn(ai_player)
		}
	}
}

func playerTurn(player *player.Player) {
	fmt.Println("It is your turn!")
	for {
		if *getInputBet() > player.GetMoney() {
			fmt.Println("You don't have that much!")
			continue
		} else {
			break
		}
	}
}

func getInputBet() *int {
	var bet int
	fmt.Println("Enter the amount you want to bet:")
	fmt.Scan(&bet)
	return &bet
}

func aiTurn(player *player.Player) {

}
