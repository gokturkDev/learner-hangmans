package main

import (
	"fmt"
	"game_utils/coin"
	"game_utils/player"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	fmt.Print("Welcome to the GOLang coin flipper-uh!\nPress enter to continue:")
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	rand.NewSource(time.Now().Unix())

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
		bet := getPlayerBet(player)
		coin := playerHeadsOrTails()
	}
}

func getPlayerBet(player *player.Player) int {
	bet := getInputBet()
	if *bet > player.GetMoney() {
		fmt.Println("You don't have that much!")
		return getPlayerBet(player)
	} else {
		return *bet
	}
}

func playerHeadsOrTails() coin.Coin {
	for {
		var selection string
		fmt.Println("Heads Or Tails? (h/t)")
		fmt.Scan(&selection)
		if selection == "h" {
			return coin.Coin{Heads: true}
		} else if selection == "t" {
			return coin.Coin{Tails: true}
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

func getAIBet(player *player.Player) int {
	return rand.Intn(player.GetMoney())
}
