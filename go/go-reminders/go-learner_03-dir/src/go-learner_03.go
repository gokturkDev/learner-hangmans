package main

import (
	"fmt"
	"game_utils/game_loop"
	"os"
	"os/exec"
)

func main() {
	fmt.Print("Welcome to the GOLang coin flipper-uh!\nPress enter to continue:")
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	game_loop.StartGameLoop()
}
