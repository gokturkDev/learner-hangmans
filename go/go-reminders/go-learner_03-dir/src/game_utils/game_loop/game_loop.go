package main

import (
	"fmt"
	"game_utils/coin"
	"game_utils/player"
	"math/rand"
	"os"
	"os/exec"
	"time"
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
	isGameOver := false
	for !isGameOver { //this is the while of GO

		fmt.Printf("Your coins left: %d \n", human_player.GetMoney())
		fmt.Printf("AI coins left: %d \n", ai_player.GetMoney())
		if isPlayerTurn {
			isGameOver = playerTurn(human_player, ai_player)
			isPlayerTurn = false
		} else {
			isGameOver = aiTurn(human_player, ai_player)
			isPlayerTurn = true
		}
	}
	fmt.Println("Thank you for playing")
}

func playerTurn(human_player, ai_player *player.Player) bool {
	fmt.Println("It is your turn!")
	bet := getPlayerBet(human_player)
	human_coin := playerHeadsOrTails()
	ai_coin := human_coin.GetOppositeCoin()
	result := evaluateResult(&human_coin, &ai_coin)
	return concludeTurn(human_player, ai_player, &result, bet)
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
			return coin.NewHeadsCoin()
		} else if selection == "t" {
			return coin.NewTailsCoin()
		}
	}
}

func getInputBet() *int {
	var bet int
	fmt.Println("Enter the amount you want to bet:")
	fmt.Scan(&bet)
	return &bet
}

func aiTurn(human_player, ai_player *player.Player) bool {
	fmt.Println("It is ai turn!")
	bet := getAIBet(ai_player)
	printAndSleep("Ai is choosing a bet...")
	fmt.Printf("AI has chosen %d for bet\n", bet)
	ai_coin := getRandCoin()
	printAndSleep(fmt.Sprint("AI has chosen %s", coinToString(&ai_coin)))
	human_coin := ai_coin.GetOppositeCoin()
	result := evaluateResult(&human_coin, &ai_coin)
	return concludeTurn(human_player, ai_player, &result, bet)
}

func getAIBet(player *player.Player) int {
	return rand.Intn(player.GetMoney())
}

func getRandCoin() coin.Coin {
	randNum := rand.Intn(2)
	if randNum == 0 {
		return coin.NewHeadsCoin()
	}
	return coin.NewTailsCoin()
}

func isTwoCoinsEqual(firstCoin, secondCoin *coin.Coin) bool {
	return firstCoin.IsHeads() == secondCoin.IsHeads()
}

type WinResult struct {
	player_won bool
	ai_won     bool
}

func whoWon(human_coin, ai_coin, winning_coin *coin.Coin) WinResult {
	if isTwoCoinsEqual(human_coin, winning_coin) {
		return WinResult{player_won: true}
	}
	return WinResult{ai_won: true}
}

func evaluateResult(human_coin, ai_coin *coin.Coin) WinResult {
	winning_coin := getRandCoin()
	win_result := whoWon(human_coin, ai_coin, &winning_coin)
	printAndSleep(fmt.Sprintf("The coin landed on %s", coinToString(&winning_coin)))
	return win_result
}

func concludeTurn(human_player, ai_player *player.Player, winResult *WinResult, bet int) bool {
	if winResult.player_won {
		return playerWonTheRound(human_player, ai_player, bet)
	}
	return aiWonTheRound(human_player, ai_player, bet)

}

func playerWonTheRound(human_player, ai_player *player.Player, bet int) bool {
	didLose := didPlayerLost(ai_player, bet)
	if didLose {
		printAndSleep("You win, ai has gone bankrupt")
		return true
	} else {
		printAndSleep("You win the round")
		ai_player.Lost(bet)
		human_player.Won(bet)
		return false
	}
}

func aiWonTheRound(human_player, ai_player *player.Player, bet int) bool {
	didLose := didPlayerLost(human_player, bet)
	if didLose {
		printAndSleep("You lost, you do not have sufficient credit to pay off the bet")

		return true
	} else {
		printAndSleep("You lost the round")
		ai_player.Won(bet)
		human_player.Lost(bet)
		return false
	}
}

func didPlayerLost(player *player.Player, bet int) bool {
	return player.GetMoney() <= bet
}

func coinToString(coin *coin.Coin) string {
	var sideInfo string
	if coin.IsHeads() {
		sideInfo = "Heads"
	} else {
		sideInfo = "Tails"
	}
	return sideInfo
}

func printAndSleep(text string) {
	fmt.Println(text)
	time.Sleep(time.Second)
}
