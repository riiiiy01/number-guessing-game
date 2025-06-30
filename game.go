package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Difficult struct {
	difficult string
	chance    uint8
}

func playGame() {
	var selectDifficult []*Difficult

	easy := &Difficult{"Easy", 10}
	medium := &Difficult{"Medium", 5}
	hard := &Difficult{"Hard", 3}

	selectDifficult = append(selectDifficult, easy, medium, hard)

	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	var attempt uint8
	var playerChoice uint8
	var playerGuess uint8

	rightAnswer := rng.Intn(100) + 1

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")

	fmt.Println("Please select the difficulty level: ")

	for i, diff := range selectDifficult {
		fmt.Printf("%d.%s (%d chances) \n", i+1, diff.difficult, diff.chance)
	}

	for {
		fmt.Printf("Enter your choice: ")
		fmt.Scanln(&playerChoice)

		if playerChoice < 1 || playerChoice > 3 {
			fmt.Println("Enter the right number!")
		} else {
			break
		}

	}

	selected := selectDifficult[playerChoice-1]

	fmt.Printf("Great! You have selected the %s difficulty level.\n", selected.difficult)

	fmt.Println("Let's start the game!")

	for {
		fmt.Printf("Enter your guess: ")
		fmt.Scanln(&playerGuess)

		if playerGuess < uint8(rightAnswer) {
			fmt.Printf("Incorrect! the number is greater than %d.\n", playerGuess)
			selected.chance--
			attempt += 1

			if selected.chance == 0 {
				fmt.Println("Game over!")
				break
			}

		} else if playerGuess > uint8(rightAnswer) {
			fmt.Printf("Incorrect! the number is less than %d.\n", playerGuess)
			selected.chance--
			attempt += 1

			if selected.chance == 0 {
				fmt.Println("Game over!")
				break
			}

		} else {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempt\n", attempt)
			break
		}

	}
}
