package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"reflect"
)

type Game struct {
	AllowedCommands map[string]string
}

// Throw returns random guess made by the program.
func (g Game) Throw() string {
	keys := reflect.ValueOf(g.AllowedCommands).MapKeys()

	return keys[rand.Intn(len(keys))].Interface().(string)
}

// Play starts the game, prompts for input, and keeps the game alive.
func (g Game) Play() {
	fmt.Print("\nType your guess (r=rock, p=paper, s=scissors): ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		playerThrown := scanner.Text()
		playerDefeatsThrow := g.AllowedCommands[playerThrown]

		if playerThrown == "" {
			fmt.Println("Thank you for playing!")
			os.Exit(0)
		}

		if playerDefeatsThrow == "" {
			fmt.Print("\nTry again (r=rock, p=paper, s=scissors): ")
		} else {
			// Valid command thrown by player
			computerThrown := g.Throw()
			computerDefeatsThrow := g.AllowedCommands[computerThrown]

			if playerDefeatsThrow == computerThrown {
				fmt.Print("You win ", "👍")
			}

			if computerDefeatsThrow == playerThrown {
				fmt.Print("You lose ", "👎")
			}

			if playerThrown == computerThrown {
				fmt.Print("Tie ", "🤝")
			}

			fmt.Printf(" → (%s) vs (%s)", playerThrown, computerThrown)
			fmt.Print("\n\nPlay again (r=rock, p=paper, s=scissors): ")
		}
	}
}
