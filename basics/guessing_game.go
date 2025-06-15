package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the guessing game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it its?")

	var guess int
	for {
		fmt.Println("Please enter a number: ")
		fmt.Scanln(&guess)

		// Check if the target is correct
		if guess == target {
			fmt.Println("You guessed the number!")
			break
		} else if guess < target {
			fmt.Println("Number is too low!")
		} else if guess > target {
			fmt.Println("Number is too high!")
		} else {
			fmt.Println("Invalid guess!")
		}
	}
}
