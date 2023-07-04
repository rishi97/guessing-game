package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a random number between 1 and 100
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	// Initialize variables
	guess := 0
	attempts := 0

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	// Loop until the player guesses the correct number
	for guess != target {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
		attempts++

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
		}
	}
}
