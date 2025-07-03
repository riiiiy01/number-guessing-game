package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Difficult struct {
	difficult string
	chance    uint8
	maxChance uint8
}

func playGame() {
	var selectDifficult []*Difficult

	easy := &Difficult{"Easy", 10, 10}
	medium := &Difficult{"Medium", 5, 5}
	hard := &Difficult{"Hard", 3, 3}

	selectDifficult = append(selectDifficult, easy, medium, hard)

	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	var attempt uint8
	var playerChoice uint8
	var playerGuess uint8

	maxScore := 100
	rightAnswer := rng.Intn(100) + 1

	start := time.Now()

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")

	fmt.Println("Please select the difficulty level: ")

	for i, diff := range selectDifficult {
		fmt.Printf("%d.%s (%d chances) \n", i+1, diff.difficult, diff.chance)
	}

	str := string(strconv.Itoa(rightAnswer))

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

	gameLogic(playerGuess, rightAnswer, attempt, start, str, maxScore, selected)
}

func gameLogic(playerGuess uint8, rightAnswer int, attempt uint8, start time.Time, hint string, maxScore int, selected *Difficult) {
	for {
		fmt.Printf("Enter your guess: ")
		fmt.Scanln(&playerGuess)

		if playerGuess < uint8(rightAnswer) {
			fmt.Printf("Incorrect! the number is greater than %d.\n", playerGuess)
		} else if playerGuess > uint8(rightAnswer) {
			fmt.Printf("Incorrect! the number is less than %d.\n", playerGuess)
		}

		// hint
		if selected.difficult == "Easy" && attempt == 6 {
			fmt.Println("Hint! the last number of number guess is", string(hint[1]))
		} else if selected.difficult == "Medium" && attempt == 2 {
			fmt.Println("Hint! the last number of number guess is", string(hint[1]))
		} else if selected.difficult == "Hard" && attempt == 1 {
			fmt.Println("Hint! the last number of number guess is", string(hint[1]))
		}

		attempt++
		selected.chance--

		// score
		var countScore int
		var score int
		countScore = maxScore - ((int(attempt) - 1) * (maxScore / int(selected.maxChance)))

		score = max(countScore, 10)

		if playerGuess == uint8(rightAnswer) {
			elapsed := time.Since(start)
			minutes := int(elapsed.Minutes())
			seconds := int(elapsed.Seconds()) % 60
			fmt.Printf("Congratulations! You guessed the correct number in %d attempt\n", attempt)
			fmt.Printf("You take %d minutes and %d seconds to guess the number \n", minutes, seconds)
			fmt.Printf("you got score: %d \n", score)
			break
		}

		if selected.chance == 0 {
			fmt.Println("Game Over!")
			fmt.Printf("you got score: %d \n", score)
			break
		}
	}
}
