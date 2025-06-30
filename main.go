package main

import (
	"fmt"
)

func main() {
	for {

		playGame()

		var playAgain string
		fmt.Printf("Do you want to play again? ")
		fmt.Scanln(&playAgain)

		if playAgain != "y" {
			break
		}

	}
}
