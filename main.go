package main

import "fmt"

type Difficult struct {
	difficult string
	chance    int
}

func main() {
	var selectDifficult []Difficult

	easy := Difficult{"Easy", 10}
	medium := Difficult{"Medium", 5}
	hard := Difficult{"Hard", 3}

	selectDifficult = append(selectDifficult, easy, medium, hard)

	var playerChoice uint8

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")

	fmt.Println("Please select the difficulty level: ")

	for i, diff := range selectDifficult {
		fmt.Printf("%d. %s (%d chances) \n", i+1, diff.difficult, diff.chance)
	}

	fmt.Printf("Enter your choice: ")
	fmt.Scanln(&playerChoice)
	fmt.Printf("Great! You have selected the %s difficulty level.\n", selectDifficult[playerChoice-1].difficult)
	fmt.Println("Let's start the game!")
}
