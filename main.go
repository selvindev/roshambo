package main

import (
	"github.com/selvindev/roshambo/game"
)

func main() {
	g := game.Game{
		AllowedCommands: map[string]string{
			"r": "s",
			"p": "r",
			"s": "p",
		},
	}

	g.Play()
}
