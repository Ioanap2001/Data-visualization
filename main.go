package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"git.sr.ht/~sbinet/gg"
)

// GameStats keeps track of game statistics
type GameStats struct {
	Played int
	Won    int
}

func main() {
	// Display message on startup
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Guess The Number")
	fmt.Println("Let's see how good you are at this")
	fmt.Println("[!] You have 5 tries, Good Luck!!!")
	fmt.Println(strings.Repeat("-", 40))

	var stats GameStats
	rand.Seed(time.Now().UnixNano())

	for {
		guess, prompt := pickLevel()
		win := start(guess, prompt)
		stats.Played++

		if win {
			stats.Won++
		}

		printResults(stats)
		err := saveGameStatus(stats)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println()
		fmt.Print("Do you want to play again? [1/0]: ")
		option := getOption()
		if option != 1 {
			break
		}

		fmt.Println()
	}

	fmt.Println("Thanks for playing!")
}

// pickLevel displays the options menu and returns the selected level's guess and prompt
func pickLevel() (int, string) {
	var guess int
	var prompt string

	// Display options menu
	fmt.Println()
	fmt.Println("Choose a level")
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("1. Easy [0 - 10]")
	fmt.Println("2. Medium [0 - 25]")
	fmt.Println("3. Hard [0 - 50]")
	fmt.Println("4. Expert [0 - 100]")
	fmt.Println("5. Legendary [0 - 1000]")
	fmt.Println(strings.Repeat("-", 25))

	// Get input
	fmt.Print("\noption: ")
	level := getOption()

	// Get the random number to guess and its corresponding prompt
	switch level {
	case 1:
		guess = rand.Intn(11)
		prompt = "Guess a number between 0 and 10: "
	case 2:
		guess = rand.Intn(26)
		prompt = "Guess a number between 0 and 25: "
	case 3:
		guess = rand.Intn(51)
		prompt = "Guess a number between 0 and 50: "
	case 4:
		guess = rand.Intn(101)
		prompt = "Guess a number between 0 and 100: "
	case 5:
		guess = rand.Intn(1001)
		prompt = "Guess a number between 0 and 1000: "
	}

	return guess, prompt
}

// start runs the guessing game
func start(guess int, prompt string) bool {
	for tries := 1; tries <= 5; tries++ {
		fmt.Print(prompt)
		userGuess := getOption()

		switch {
		case userGuess == guess:
			fmt.Println("You guessed correctly, Good job!")
			if tries == 1 {
				fmt.Println("Wow! You are amazing.")
				fmt.Println("You guessed the number correctly on your first attempt!")
			} else {
				fmt.Printf("It took you %d tries\n", tries)
			}
			return true
		case tries < 5:
			if userGuess < guess {
				fmt.Println("Too low, try again!")
			} else if userGuess > guess {
				fmt.Println("Too high, try again!")
			}
		}
	}

	fmt.Println("You lost, The number is", guess)
	fmt.Println("Better luck next time.")
	return false
}

// getOption gets an integer option from the user
func getOption() int {
	var option int
	_, err := fmt.Scan(&option)
	if err != nil {
		log.Panicln("Please enter an integer")
	}
	return option
}

// printResults prints the game results
func printResults(stats GameStats) {
	fmt.Println()
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("Results")
	fmt.Println(strings.Repeat("-", 25))
	fmt.Printf("You played %d game(s).\n", stats.Played)
	fmt.Printf("You won %d game(s).\n", stats.Won)
}

// saveGameStatus saves the game status as a PNG graphic
func saveGameStatus(stats GameStats) error {
	const width = 400
	const height = 200

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	// Draw text
	dc.DrawString(fmt.Sprintf("Games Played: %d", stats.Played), 20, 40)
	dc.DrawString(fmt.Sprintf("Games Won: %d", stats.Won), 20, 80)

	// Save the image
	if err := dc.SavePNG("game_status.png"); err != nil {
		return err
	}
	fmt.Println("Game status saved as 'game_status.png'")

	// Generate and save the pie chart
	err := GeneratePieChart(stats.Played, stats.Won)
	if err != nil {
		return err
	}
	return nil
}
